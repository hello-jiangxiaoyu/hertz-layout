package cmd

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz/demo/biz/dal"
	"os"
	"time"
)

// StartAdminServer 管理面服务接口
func StartAdminServer() {
	var err error
	for i := 0; i < 3; i++ {
		if err = dal.InitAdmin(); err == nil {
			break
		}
		hlog.Error("### admin init dal err: ", err.Error())
		time.Sleep(10 * time.Second)
	}
	if err != nil {
		os.Exit(1)
		return
	}

	h := server.Default(server.WithHostPorts(":1900"))
	h.Spin()
}
