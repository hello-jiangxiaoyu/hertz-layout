package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"hertz/demo/internal/conf"
	"hertz/demo/internal/response"
	"net"
	"os"
	"runtime"
	"strings"
)

func stackInfo(msg string, err any, skip int, fullStack bool) string {
	pwd, _ := os.Getwd()
	pwd = strings.ReplaceAll(pwd, `\`, "/") // handle windows path
	res := fmt.Sprintf("panic recovered: %s; %v", err, msg)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		if fullStack || pwd == "" || strings.Contains(file, pwd) { // only about quick auth source files
			res += fmt.Sprintf("\n\t%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
		}
	}
	return res + "\n\n"
}

// MyRecoveryHandler panic处理
func MyRecoveryHandler(ctx context.Context, c *app.RequestContext) {
	defer func() {
		if err := recover(); err != nil {
			// Check for a broken connection, as it is not really a
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				var se *os.SyscallError
				if errors.As(ne, &se) {
					seStr := strings.ToLower(se.Error())
					if strings.Contains(seStr, "broken pipe") || strings.Contains(seStr, "connection reset by peer") {
						brokenPipe = true
					}
				}
			}
			req := fmt.Sprintf("method:%s path:%s", c.Method(), c.Path())
			hlog.Error(stackInfo(req, err, 3, conf.GetConf().Hertz.LogFullStack))
			if !brokenPipe {
				response.ErrorPanic(c)
			}
		}
	}()
	c.Next(ctx)
}
