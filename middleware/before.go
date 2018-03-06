package middleware

import (
	"encoding/json"

	"github.com/kataras/iris"
)

// Before is
func Before(ctx iris.Context) {

	if ctx.Method() == "OPTIONS" {
		ctx.WriteString("ok")
		return
	}
	// shareInformation := "this is a sharable information between handlers"
	// ctx.Values().Set("info", shareInformation)
	u := ctx.Request().UserAgent()
	ip := "\nip:" + ctx.RemoteAddr()
	// for the sake of simplicity, in order see the logs at the ./_today_.txt
	fmsMap := ctx.FormValues()
	fmsJSON, _ := json.Marshal(fmsMap)

	fmsStr := "\nFms:" + string(fmsJSON)
	info := "Request path:" + ctx.Path() + ip + fmsStr + "\nUserAgent: " + u

	ctx.Application().Logger().Info(info)
	ctx.Next()
}

// After is
func After(ctx iris.Context) {
	println("After the indexHandler or contactHandler")
}
