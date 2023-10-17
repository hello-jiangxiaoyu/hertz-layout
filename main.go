// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz/demo/biz/dal"
	"os"
	"time"
)

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

	h := server.Default()
	register(h)
	h.Spin()
}
