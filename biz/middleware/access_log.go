package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertz/demo/biz/dal/log"
	"time"
)

func AccessLog(ctx context.Context, c *app.RequestContext) {
	start := time.Now()
	c.Next(ctx)
	log.Access.Info().
		Int("status", c.Response.StatusCode()).
		Str("method", string(c.Method())).
		Str("path", string(c.Path())).
		Str("full_path", c.FullPath()).
		Str("host", string(c.Host())).
		Str("client_ip", c.ClientIP()).
		Int64("latency", start.Sub(time.Now()).Milliseconds()).
		Int("byte_received", len(c.Request.Body())).
		Int("byte_sent", len(c.Response.Body())).
		Str("proto", string(c.Request.URI().Scheme())).
		Str("ref", c.Request.Header.Get(consts.HeaderReferer)).
		Str("ua", c.Request.Header.Get(consts.HeaderUserAgent)).
		Msg(c.Errors.String())
}
