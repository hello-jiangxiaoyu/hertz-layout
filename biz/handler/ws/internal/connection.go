package internal

import (
	"github.com/bytedance/sonic"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"hertz/demo/pkg/utils"
)

var (
	Sender            WebSocketConnection
	ErrorUserNotExist = errors.New("user not exist")
)

type WebSocketConnection struct {
	Connection utils.MapInt64[websocket.Conn]
	RoomUserID utils.MapInt64[[]int64]
}

type Message struct {
	ReceiveID int64
	Path      string
	Content   string
}

type MessageSend struct {
	SenderID int64
	RoomID   int64
	Type     string
	Content  string
}

func (c *WebSocketConnection) SendUserMessage(senderID int64, userID int64, roomID int64, msg *Message) error {
	conn := c.Connection.Get(userID)
	if conn == nil {
		return ErrorUserNotExist
	}

	msgSend := MessageSend{
		SenderID: senderID,
		RoomID:   roomID,
		Type:     msg.Path,
		Content:  msg.Content,
	}
	byteData, err := sonic.Marshal(&msgSend)
	if err != nil {
		return err
	}
	if err = conn.WriteMessage(websocket.TextMessage, byteData); err != nil {
		return errors.WithMessage(err, "conn.WriteMessage err")
	}
	return nil
}
