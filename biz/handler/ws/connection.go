package ws

import (
	"github.com/bytedance/sonic"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"hertz/demo/pkg/utils"
)

var (
	ErrorUserNotExist = errors.New("user not exist")
	sender            WebSocketConnection
)

type WebSocketConnection struct {
	Connection utils.MapInt64[websocket.Conn]
	RoomUserID utils.MapInt64[[]int64]
}

type Message struct {
	ReceiveID int64
	Type      string
	Content   string
}

type MessageSend struct {
	SenderID int64
	RoomID   int64
	Type     string
	Content  string
}

func (c *WebSocketConnection) SendRoomMessage(senderID int64, roomID int64, msg *Message) error {
	userIDs := c.RoomUserID.Get(roomID)
	if userIDs == nil {
		return errors.New("room not exist")
	}
	for _, userID := range *userIDs {
		if err := c.SendUserMessage(senderID, userID, roomID, msg); err != nil && !errors.Is(err, ErrorUserNotExist) {
			return err
		}
	}
	return nil
}

func (c *WebSocketConnection) SendUserMessage(senderID int64, userID int64, roomID int64, msg *Message) error {
	conn := c.Connection.Get(userID)
	if conn == nil {
		return ErrorUserNotExist
	}

	msgSend := MessageSend{
		SenderID: senderID,
		RoomID:   roomID,
		Type:     msg.Type,
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
