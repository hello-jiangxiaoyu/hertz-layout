// Code generated by hertz generator.

package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"hertz/demo/biz/handler/internal"
	"hertz/demo/pkg/response"
	"time"
)

const (
	pongWait   = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.

	MsgTypeUser = "user" // 好友消息
	MsgTypeRoom = "room" //	群消息
)

var (
	a        internal.Api
	upgrader = websocket.HertzUpgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(c *app.RequestContext) bool {
			return true
		},
	}
)

// ChatRoom .
// @Description	websocket 聊天api
// @Tags		ws
// @Success		200
// @router		/v1/hertz/ws [GET]
func ChatRoom(_ context.Context, c *app.RequestContext) {
	if err := a.SetReqWithSub(c).Error; err != nil {
		response.ErrorRequest(c, err)
		return
	}

	if err := upgrader.Upgrade(c, WebsocketHandler(*a.Sub)); err != nil {
		response.ErrorUnknown(c, err, "websocket upgrade err")
		return
	}
	response.Success(c)
}