package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/handler/admin"
	"anla.io/taizhou-fe-api/handler/user"
	"anla.io/taizhou-fe-api/middleware"
	"github.com/kataras/iris"
)

// AdminRouter is
func AdminRouter(adminParty iris.Party) {
	ad := adminParty.Party("/admin")
	{
		ad.Post("/login", user.Login)
	}

	adAuth := ad.Party("/a", jwt.JwtHandler.Serve)
	{
		adAuth.Get("/user", user.User{}.GetAll)
		adAuth.Post("/user/logout", user.User{}.GetMe)
		adAuth.Post("/category", middleware.Edit, admin.Category{}.Create)
		adAuth.Put("/category", middleware.Edit, admin.Category{}.Update)
		adAuth.Post("/register", middleware.Edit, handler.Register{}.AdminAdd)
	}
}
