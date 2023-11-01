package internal

import (
	"github.com/bytedance/sonic"
	"github.com/hertz-contrib/websocket"
)

type WsEngine struct {
	handler map[string]func(*websocket.Conn, int64, *Message) error
}

func (e *WsEngine) AddHandler(path string, handler func(conn *websocket.Conn, sub int64, msg *Message) error) {
	e.handler[path] = handler
}

func (e *WsEngine) Handle(conn *websocket.Conn, sub int64, msgByte []byte) error {
	var msg Message
	if err := sonic.Unmarshal(msgByte, &msg); err != nil {
		return ErrorUnknown(conn, err, "unmarshal text message err")
	}

	handler, ok := e.handler[msg.Path]
	if !ok {
		return ErrorMsg(conn, "no such path")
	}
	return handler(conn, sub, &msg)
}
