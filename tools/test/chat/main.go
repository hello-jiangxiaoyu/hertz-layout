package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/websocket"
)

var upgrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

// websocket 广播聊天室案例
func main() {
	go hub.run()
	h := server.Default(server.WithHostPorts(":8080"))
	h.GET("/ws", func(c context.Context, ctx *app.RequestContext) {
		serveWs(ctx)
	})

	h.Spin()
}
