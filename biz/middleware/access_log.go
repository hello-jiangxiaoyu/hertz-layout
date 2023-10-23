package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type options struct {
	format           string
	timeFormat       string
	timeInterval     time.Duration
	logFunc          func(ctx context.Context, format string, v ...interface{})
	timeZoneLocation *time.Location
	enableLatency    bool
}

var defaultTagFormat = "[${time}] ${status} - ${latency} ${method} ${path}"
var defaultFormat = " %s | %3d | %7v | %-7s | %-s "

func AccessLog(ctx context.Context, c *app.RequestContext) {

}
