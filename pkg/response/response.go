package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

// Response 结构体类型响应
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

// ArrayResponse 数组类型响应
type ArrayResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Total int    `json:"total,omitempty"`
	Data  any    `json:"data,omitempty"`
}

// 构造响应
func response(c *app.RequestContext, code int, errCode int, err error, msg string, data any, total int, isArray []bool) {
	c.Header("X-Request-Id", c.GetString("requestID"))
	if len(isArray) == 0 || !isArray[0] {
		if data == nil {
			data = struct{}{}
		}
		c.JSON(code, &Response{Code: errCode, Msg: msg, Data: data})
	} else {
		if data == nil {
			data = []struct{}{}
		}
		c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})
	}

	if err != nil {
		_ = c.Error(errors.WithMessage(err, msg))
	} else {
		_ = c.Error(errors.New(msg))
	}
	c.Set("code", errCode)
	c.Abort()
}
