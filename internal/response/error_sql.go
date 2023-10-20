// sql操作错误相关响应

package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gorm.io/gorm"
	"strings"
)

// ErrorUpdate SQL更新失败
func ErrorUpdate(c *app.RequestContext, err error, respMsg string, isArray ...bool) {
	if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		errorResponse(c, consts.StatusInternalServerError, CodeSqlModifyDuplicate, err, respMsg, isArray)
	} else {
		errorResponse(c, consts.StatusInternalServerError, CodeSqlModify, err, respMsg, isArray)
	}
}

func ErrorCreate(c *app.RequestContext, err error, respMsg string, isArray ...bool) {
	if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		errorResponse(c, consts.StatusConflict, CodeSqlCreateDuplicate, err, "Duplicate field name", isArray)
	} else {
		errorResponse(c, consts.StatusInternalServerError, CodeSqlCreate, err, respMsg, isArray)
	}
}

// ErrorSelect 数据库查询错误
func ErrorSelect(c *app.RequestContext, err error, respMsg string, isArray ...bool) {
	if err != nil && strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) { // gorm find操作record not found
		errorResponse(c, consts.StatusNotFound, CodeSqlSelectNotFound, err, respMsg, isArray)
	} else {
		errorResponse(c, consts.StatusInternalServerError, CodeSqlSelect, err, respMsg, isArray)
	}
}

// ErrorDelete 数据库删除错误
func ErrorDelete(c *app.RequestContext, err error, respMsg string, isArray ...bool) {
	if err != nil && strings.Contains(err.Error(), gorm.ErrForeignKeyViolated.Error()) { // 外键依赖导致无法删除
		errorResponse(c, consts.StatusConflict, CodeSqlDeleteForKey, err, respMsg, isArray)
	} else {
		errorResponse(c, consts.StatusInternalServerError, CodeSqlDelete, err, respMsg, isArray)
	}
}
