// Package resp 未知错误响应
package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ErrorUnknown 未知错误
func ErrorUnknown(c *app.RequestContext, err error, respMsg string, isArray ...bool) {
	errorResponse(c, consts.StatusInternalServerError, CodeUnknown, err, respMsg, isArray)
}

// ErrorNotFound 资源未找到
func ErrorNotFound(c *app.RequestContext, err error, respMsg string, isArray ...bool) {
	errorResponse(c, consts.StatusInternalServerError, CodeNotFound, err, respMsg, isArray)
}

func ErrorSaveSession(c *app.RequestContext, err error, isArray ...bool) {
	errorResponse(c, consts.StatusInternalServerError, CodeSaveSession, err, "failed to save session", isArray)
}
