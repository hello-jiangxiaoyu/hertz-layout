// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"hertz/demo/biz/dal"
	"hertz/demo/biz/dal/log"
	"hertz/demo/biz/middleware"
	"hertz/demo/internal/conf"
	"os"
	"time"
)

// @title HertzTest
// @version 1.0
// @description		This is a demo using Hertz.
// @license.name	Apache 2.0
// @host			localhost:8888
// @BasePath /
// @schemes http
func main() {
	var err error
	for i := 0; i < 3; i++ {
		if err = dal.Init(); err == nil {
			break
		}
		hlog.Error("### init dal err: ", err.Error())
		time.Sleep(10 * time.Second)
	}
	if err != nil {
		os.Exit(1)
		return
	}

	log.Zero.Info().Msg("zero log")
	log.Access.Info().Msg("access log")

	hz := getServer()
	hz.Spin()
}

func getServer() *server.Hertz {
	opts := []config.Option{
		server.WithHostPorts(conf.GetConf().Hertz.Address),                 // 监听端口
		server.WithTracer(middleware.PrometheusHandler(":9091", "/hertz")), // 请求量和时延等监控
	}
	h := server.New(opts...)
	h.Use(middleware.MyRecoveryHandler)       // panic处理
	h.Use(gzip.Gzip(gzip.DefaultCompression)) //响应压缩
	h.Use(accesslog.New())
	h.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusBadRequest, "no such router")
	})
	register(h)
	return h
}
