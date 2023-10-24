package cmd

import (
	"context"
	"os"
	"time"

	"hertz/demo/biz/dal"
	"hertz/demo/biz/handler"
	"hertz/demo/biz/middleware"
	"hertz/demo/biz/router"
	"hertz/demo/pkg/conf"
	_ "hertz/demo/pkg/docs"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/swagger"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
)

// StartServer 数据面服务接口
func StartServer(_ *cobra.Command, _ []string) {
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

	hz := getServer()
	hz.Spin()
}

// 初始化hertz服务
func getServer() *server.Hertz {
	opts := []config.Option{
		server.WithHostPorts(conf.GetConf().Hertz.Address),                 // 监听端口
		server.WithTracer(middleware.PrometheusHandler(":9091", "/hertz")), // 请求量和时延等监控
	}
	h := server.New(opts...)
	h.Use(middleware.MyRecoveryHandler)       // panic处理
	h.Use(middleware.AccessLog)               // 请求日志
	h.Use(gzip.Gzip(gzip.DefaultCompression)) // 响应压缩
	h.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusBadRequest, "no such router")
	})
	router.GeneratedRegister(h) // 服务接口
	customizedRegister(h)
	return h
}

// 自定义接口
func customizedRegister(r *server.Hertz) {
	r.GET("/v1/hertz/ping", handler.Ping)
	r.GET("/v1/hertz/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
}
