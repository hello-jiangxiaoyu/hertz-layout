package internal

import (
	"github.com/bytedance/sonic"
	"github.com/hertz-contrib/websocket"
	"hertz/demo/biz/dal/log"
	"time"
)

const (
	maxMessageSize = 1024                // Maximum message size allowed from peer.
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
)

type WsEngine struct {
	handler map[string]func(ctx *WsContext) error
	sender  *WebSocketConnection
}

type WsContext struct {
	Conn   *websocket.Conn
	Sender *WebSocketConnection
	Msg    *Message
	Sub    int64
}

func NewWsEngine() *WsEngine {
	return &WsEngine{
		handler: make(map[string]func(ctx *WsContext) error),
		sender:  &WebSocketConnection{},
	}
}

// AddHandler 添加一个消息类型处理函数
func (e *WsEngine) AddHandler(path string, handler func(ctx *WsContext) error) {
	e.handler[path] = handler
}

// WebsocketHandler 返回websocket处理函数
func (e *WsEngine) WebsocketHandler(sub int64) websocket.HertzHandler {
	return func(conn *websocket.Conn) {
		e.initConnection(sub, conn)
		ticker := time.NewTicker(pingPeriod)
		defer func() {
			ticker.Stop()
			e.sender.Connection.Delete(sub) // 用户下线
			if err := conn.Close(); err != nil {
				log.Error().Msgf("close websocket conn err: %s", err.Error())
			}
		}()

		// 消息接收处理
		for {
			mt, message, err := conn.ReadMessage() // 获取客户端发送的消息
			if err != nil {
				log.Error().Msg("read message err" + err.Error())
				break // 客户端关闭连接
			}

			if err = e.HandleMessage(conn, sub, mt, message); err != nil {
				return
			}
		}
	}
}

// HandleMessage 消息内容处理
func (e *WsEngine) HandleMessage(conn *websocket.Conn, sub int64, mt int, msgByte []byte) error {
	var msg Message
	if err := sonic.Unmarshal(msgByte, &msg); err != nil {
		return ErrorUnknown(conn, err, "unmarshal text message err")
	}

	switch mt {
	case websocket.PingMessage:
		if err := PongResponse(conn); err != nil {
			log.Error().Msg("conn.WriteMessage pong err" + err.Error())
		}
	case websocket.TextMessage:
		handler, ok := e.handler[msg.Path]
		if !ok {
			return ErrorMsg(conn, "no such path")
		}
		ctx := &WsContext{
			Sender: e.sender,
			Conn:   conn,
			Msg:    &msg,
			Sub:    sub,
		}
		return handler(ctx)
	case websocket.PongMessage:
		// todo: 看门狗关闭连接
	}

	return nil
}

func (e *WsEngine) initConnection(sub int64, conn *websocket.Conn) {
	conn.SetReadLimit(maxMessageSize)
	_ = conn.SetReadDeadline(time.Now().Add(pongWait))
	_ = conn.SetWriteDeadline(time.Now().Add(writeWait))
	conn.SetPongHandler(func(string) error { return conn.SetReadDeadline(time.Now().Add(pongWait)) })
	e.sender.Connection.Set(sub, conn)
}
