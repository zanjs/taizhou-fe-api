package middleware

import (
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// AdminEdit is
func AdminEdit(ctx iris.Context) {
	user := User{}.GetUser(ctx)
	ctx.Application().Logger().Info("AdminEdit middle：" + user.Username)
	if user.Role < 90 {
		response.JSONErrorCode(ctx, "您权限不够", response.ErrorCode.RoleErr)
		ctx.Application().Logger().Info(response.ErrorMessage.RoleErr)
		return
	}

	ctx.Next()
}

// Edit is
func Edit(ctx iris.Context) {
	user := User{}.GetUser(ctx)
	ctx.Application().Logger().Info("Edit middle：" + user.Username)
	if user.Role < 80 {
		response.JSONErrorCode(ctx, "您权限不够", response.ErrorCode.RoleErr)
		ctx.Application().Logger().Info(response.ErrorMessage.RoleErr)
		return
	}

	ctx.Next()
}

// UserEdit is
func UserEdit(ctx iris.Context) {
	user := User{}.GetUser(ctx)
	ctx.Application().Logger().Info("UserEdit middle：" + user.Username)
	if user.Role < 70 {
		response.JSONErrorCode(ctx, "您权限不够", response.ErrorCode.RoleErr)
		ctx.Application().Logger().Info(response.ErrorMessage.RoleErr)
		return
	}
	ctx.Next()
}
