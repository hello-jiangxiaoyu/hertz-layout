// Code generated by hertz generator.

package user

import (
	"context"
	"hertz/demo/biz/dal/mysql"
	"hertz/demo/biz/handler/internal"
	"hertz/demo/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
	user "hertz/demo/biz/model/hertz/user"
)

var a internal.Api

// GetUser .
// @Description	get user list
// @Tags		user
// @Success		200
// @router		/v1/hertz/admin/user [GET]
func GetUser(_ context.Context, c *app.RequestContext) {
	var req user.CommonGetReq
	if err := a.SetReqWithSub(c, &req).Error; err != nil {
		response.ErrorRequest(c, err)
		return
	}

	list, err := mysql.GetUserList(req.Cursor, req.Num)
	if err != nil {
		response.ErrorSelect(c, err, "get user list err")
		return
	}
	response.SuccessArrayData(c, list)
}

// CreateUser .
// @Description	create user
// @Tags		user
// @Success		200
// @router		/v1/hertz/admin/user [POST]
func CreateUser(_ context.Context, c *app.RequestContext) {
	var req user.CreateUserReq
	if err := a.SetReqWithSub(c, &req).Error; err != nil {
		response.ErrorRequest(c, err)
		return
	}

	if err := mysql.CreateUser(req.Username, req.Password); err != nil {
		response.ErrorCreate(c, err, "create user err")
		return
	}

	response.Success(c)
}

// DisableUser .
// @Description	disable user
// @Tags		user
// @Param		userID		path	integer	true	"user id"
// @Success		200
// @router		/v1/hertz/admin/user/{userID}/disable [PUT]
func DisableUser(_ context.Context, c *app.RequestContext) {
	var req user.UserReq
	if err := a.SetReqWithSub(c, &req).Error; err != nil {
		response.ErrorRequest(c, err)
		return
	}
	if err := mysql.DisableUser(req.UserID); err != nil {
		response.ErrorUpdate(c, err, "disable user err")
		return
	}
	response.Success(c)
}

// DeleteUser .
// @Description	delete user
// @Tags		user
// @Param		userID		path	integer	true	"user id"
// @Success		200
// @router		/v1/hertz/admin/user/{userID} [DELETE]
func DeleteUser(_ context.Context, c *app.RequestContext) {
	var req user.UserReq
	if err := a.SetReqWithSub(c, &req).Error; err != nil {
		response.ErrorRequest(c, err)
		return
	}
	if err := mysql.DeleteUser(req.UserID); err != nil {
		response.ErrorDelete(c, err, "delete user err")
		return
	}
	response.Success(c)
}
