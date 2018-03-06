package middleware

import (
	"github.com/kataras/iris"
)

// OptionsSuccess is
func OptionsSuccess(ctx iris.Context) {
	if ctx.Method() == "OPTIONS" {
		return
	}
	ctx.Next()
}
