package ws

import (
	"github.com/bytedance/sonic"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"hertz/demo/biz/dal/log"
	"hertz/demo/biz/handler/ws/internal"
	"time"
)

// WebsocketHandler websocket客户端连接处理
func WebsocketHandler(sub int64) websocket.HertzHandler {
	return func(conn *websocket.Conn) {
		initConnection(sub, conn)
		ticker := time.NewTicker(pingPeriod)
		defer func() {
			ticker.Stop()
			internal.Sender.Connection.Delete(sub) // 用户下线
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
				if err = internal.PongResponse(conn); err != nil {
					log.Error().Msg("conn.WriteMessage pong err" + err.Error())
				}
			case websocket.TextMessage:
				if err = HandleTextMessage(conn, sub, message); err != nil {
					log.Error().Msg("handle text message err" + err.Error())
				}
			case websocket.PongMessage:
				// todo: 看门狗关闭连接
			}
		}
	}
}

// HandleTextMessage json业务类型消息处理
func HandleTextMessage(conn *websocket.Conn, sub int64, message []byte) error {
	var msg internal.Message
	if err := sonic.Unmarshal(message, &msg); err != nil {
		return internal.ErrorUnknown(conn, err, "unmarshal text message err")
	}

	switch msg.Path {
	case MsgTypeUser:
		if err := internal.Sender.SendUserMessage(sub, msg.ReceiveID, 0, &msg); err != nil {
			return internal.ErrorUnknown(conn, err, "send user message err")
		}
	case MsgTypeRoom:
		if err := internal.Sender.SendRoomMessage(sub, msg.ReceiveID, &msg); err != nil {
			return errors.WithMessage(err, "send room message err")
		}
	default:
		return internal.ErrorMsg(conn, "text msg type unknown")
	}

	return internal.SuccessResponse(conn)
}

func initConnection(sub int64, conn *websocket.Conn) {
	const (
		maxMessageSize = 1024             // Maximum message size allowed from peer.
		writeWait      = 10 * time.Second // Time allowed to write a message to the peer.
	)
	conn.SetReadLimit(maxMessageSize)
	_ = conn.SetReadDeadline(time.Now().Add(pongWait))
	_ = conn.SetWriteDeadline(time.Now().Add(writeWait))
	conn.SetPongHandler(func(string) error { return conn.SetReadDeadline(time.Now().Add(pongWait)) })
	internal.Sender.Connection.Set(sub, conn)
}
