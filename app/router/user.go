package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/handler/user"
	"github.com/kataras/iris"
)

// UserRouter is
func UserRouter(party iris.Party) {

	party.Post("/login", user.Login)
	party.Post("/register", handler.Register{}.Add)
	party.Get("/", jwt.JwtHandler.Serve, handler.Controller{}.JWTHandler)

	o := party.Party("/user")
	{
		o.Get("/", user.User{}.GetAll)
	}
	a := o.Party("/a")
	a.Use(jwt.JwtHandler.Serve)
	{
		a.Get("/me", user.User{}.GetMe)
	}

}
