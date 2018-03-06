package middleware

import (
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

type (
	// Login is
	Login struct{}
)

// CheckUserNameAdnPWD is 验证表单完整
func CheckUserNameAdnPWD(ctx iris.Context) {
	username := ctx.FormValue("username")

	if username == "" {
		response.JSONError(ctx, "Username where?")
		return
	}

	password := ctx.FormValue("password")

	if password == "" {
		response.JSONError(ctx, "Password where?")
		return
	}

	ctx.Next()
}
