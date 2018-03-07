package handler

import (
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// OptionsHandler is
func OptionsHandler(ctx iris.Context) {
	response.JSON(ctx, "z")
}
