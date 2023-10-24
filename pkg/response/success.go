package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 成功响应
func success(c *app.RequestContext, data any, total int, isArray ...bool) {
	response(c, consts.StatusOK, CodeSuccess, nil, "", data, total, isArray)
}

func Success(c *app.RequestContext, isArray ...bool) {
	success(c, nil, 0, isArray...)
}
func SuccessWithData(c *app.RequestContext, data any) {
	success(c, data, 0, false)
}

func SuccessArrayDataWithTotal(c *app.RequestContext, total int, data any) {
	success(c, data, total, true)
}
func SuccessArrayData(c *app.RequestContext, data any) {
	success(c, data, 0, true)
}

func DoNothing(c *app.RequestContext, msg string, isArray ...bool) {
	response(c, consts.StatusAccepted, CodeAccept, nil, msg, nil, 0, isArray)
}
