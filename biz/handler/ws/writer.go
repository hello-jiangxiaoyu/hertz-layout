package ws

import (
	"github.com/hertz-contrib/websocket"
	"time"
)

func MessageWriter(ticker *time.Ticker, conn *websocket.Conn) {
	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				continue
			}
		}
	}
}
