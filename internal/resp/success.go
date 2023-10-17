package resp

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func success(c *app.RequestContext, data any, total int, isArray ...bool) {
	response(c, consts.StatusOK, CodeSuccess, nil, "", data, total, isArray)
}

func Success(c *app.RequestContext, isArray ...bool) {
	success(c, nil, 0, isArray...)
}
func SuccessWithData(c *app.RequestContext, data any) {
	success(c, data, 0, false)
}
func SuccessArrayData(c *app.RequestContext, total int, data any) {
	success(c, data, total, true)
}

func DoNothing(c *app.RequestContext, msg string, isArray ...bool) {
	response(c, consts.StatusAccepted, CodeAccept, nil, msg, nil, 0, isArray)
}
