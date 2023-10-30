package ws

import (
	"github.com/bytedance/sonic"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"hertz/demo/biz/dal/log"
	"time"
)

// WebsocketHandler websocket客户端连接处理
func WebsocketHandler(sub int64) websocket.HertzHandler {
	return func(conn *websocket.Conn) {
		initConnection(conn)
		sender.Connection.Set(sub, conn) // 用户上线

		ticker := time.NewTicker(pingPeriod)
		// go MessageWriter(ticker, conn)
		defer func() {
			ticker.Stop()
			sender.Connection.Delete(sub) // 用户下线
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

			switch mt {
			case websocket.PingMessage:
				if err = conn.WriteMessage(websocket.PongMessage, nil); err != nil {
					log.Error().Msg("conn.WriteMessage pong err" + err.Error())
				}
			case websocket.TextMessage:
				if err = HandleTextMessage(sub, message); err != nil {
					log.Error().Msg("handle text message err" + err.Error())
				}
			case websocket.PongMessage:
				// todo: 看门狗关闭连接
			}
		}
	}
}

// HandleTextMessage json业务类型消息处理
func HandleTextMessage(sub int64, message []byte) error {
	var msg Message
	if err := sonic.Unmarshal(message, &msg); err != nil {
		return errors.WithMessage(err, "unmarshal text message err")
	}

	switch msg.Type {
	case MsgTypeUser:
		if err := sender.SendUserMessage(sub, msg.ReceiveID, 0, &msg); err != nil {
			return errors.WithMessage(err, "send user message err")
		}
	case MsgTypeRoom:
		if err := sender.SendRoomMessage(sub, msg.ReceiveID, &msg); err != nil {
			return errors.WithMessage(err, "send room message err")
		}
	default:
		return errors.New("msg type unknown")
	}
	return nil
}

func initConnection(conn *websocket.Conn) {
	const (
		maxMessageSize = 1024             // Maximum message size allowed from peer.
		writeWait      = 10 * time.Second // Time allowed to write a message to the peer.
	)
	conn.SetReadLimit(maxMessageSize)
	_ = conn.SetReadDeadline(time.Now().Add(pongWait))
	_ = conn.SetWriteDeadline(time.Now().Add(writeWait))
	conn.SetPongHandler(func(string) error { return conn.SetReadDeadline(time.Now().Add(pongWait)) })
}
