package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler"
	"github.com/kataras/iris"
)

// UserRouter is
func UserRouter(party iris.Party) {
	o := party.Party("/user")
	{
		o.Get("/", handler.User{}.GetAll)
	}
	a := o.Party("/a")
	a.Use(jwt.JwtHandler.Serve)
	{
		a.Get("/me", handler.User{}.GetMe)
	}

}
