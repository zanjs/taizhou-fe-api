package handler

import (
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// JWTError is
func JWTError(ctx iris.Context, str string) {
	response.JSONError(ctx, str)
}

// Claims is
type Claims map[string]interface{}

// JWTHandler is
func (ctl Controller) JWTHandler(ctx iris.Context) {
	user := ctl.GetUser(ctx)
	response.JSON(ctx, user)
}
