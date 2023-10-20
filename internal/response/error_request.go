// 请求错误相关响应

package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func errorResponse(c *app.RequestContext, code int, errCode int, err error, msg string, isArray []bool) {
	response(c, code, errCode, err, msg, nil, 0, isArray)
}

// ErrorRequest 请求参数错误
func ErrorRequest(c *app.RequestContext, err error, isArray ...bool) {
	errorResponse(c, consts.StatusBadRequest, CodeRequestPara, err, "invalidate request para", isArray)
}

// ErrorRequestWithMsg 请求参数错误
func ErrorRequestWithMsg(c *app.RequestContext, err error, msg string, isArray ...bool) {
	errorResponse(c, consts.StatusBadRequest, CodeRequestPara, err, msg, isArray)
}

// ErrorForbidden 无权访问
func ErrorForbidden(c *app.RequestContext, msg string, isArray ...bool) {
	errorResponse(c, consts.StatusForbidden, CodeForbidden, nil, msg, isArray)
}

// ErrorInvalidateToken token 无效
func ErrorInvalidateToken(c *app.RequestContext, isArray ...bool) {
	errorResponse(c, consts.StatusForbidden, CodeInvalidToken, nil, "invalidated token", isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(c *app.RequestContext, isArray ...bool) {
	errorResponse(c, consts.StatusUnauthorized, CodeNotLogin, nil, "user not login", isArray)
}
