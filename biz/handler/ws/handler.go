package ws

import (
	"hertz/demo/biz/handler/ws/internal"
)

var eg *internal.WsEngine

func init() {
	eg = internal.NewWsEngine()
	eg.AddHandler("/user", UserMessageHandler)
	eg.AddHandler("/group", GroupMessageHandler)
}

func UserMessageHandler(c *internal.WsContext) error {
	return nil
}

func GroupMessageHandler(c *internal.WsContext) error {
	return nil
}
