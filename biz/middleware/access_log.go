package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertz/demo/biz/dal/log"
	"time"
)

const (
	LabelTime         = "ts"
	LabelStatus       = "status"
	LabelMethod       = "method"
	LabelPath         = "path"
	LabelFullPath     = "full_path"
	LabelHost         = "host"
	LabelClientIP     = "client_ip"
	LabelLatency      = "latency"
	LabelByteReceived = "byte_received"
	LabelByteSent     = "byte_sent"
	LabelProto        = "proto"
	LabelRef          = "ref"
	LabelUA           = "ua"
	LabelErr          = "err"
)

func AccessLog(ctx context.Context, c *app.RequestContext) {
	start := time.Now()
	c.Next(ctx)
	log.Access.Info().
		Str(LabelTime, time.Now().Format("2006-01-02T15:04:05.000")).
		Int(LabelStatus, c.Response.StatusCode()).
		Str(LabelMethod, string(c.Method())).
		Str(LabelPath, string(c.Path())).
		Str(LabelFullPath, c.FullPath()).
		Str(LabelHost, string(c.Host())).
		Str(LabelClientIP, c.ClientIP()).
		Int64(LabelLatency, time.Now().Sub(start).Milliseconds()).
		Int(LabelByteReceived, len(c.Request.Body())).
		Int(LabelByteSent, len(c.Response.Body())).
		Str(LabelProto, string(c.Request.URI().Scheme())).
		Str(LabelRef, c.Request.Header.Get(consts.HeaderReferer)).
		Str(LabelUA, c.Request.Header.Get(consts.HeaderUserAgent)).
		Str(LabelErr, c.Errors.String()).Msg("")
}
