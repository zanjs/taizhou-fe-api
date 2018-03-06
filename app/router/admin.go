package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler/admin"
	"github.com/kataras/iris"
)

// AdminRouter is
func AdminRouter(adminParty iris.Party) {
	ad := adminParty.Party("/admin")
	{
		ad.Post("/register", admin.Register{}.Add)
		ad.Post("/login", admin.Login)
	}

	adAuth := ad.Party("/a", jwt.JwtHandlerAdmin.Serve)
	{
		adAuth.Get("/user/info", admin.User{}.GetMe)
		adAuth.Get("/user", admin.User{}.GetAll)
		adAuth.Post("/user/logout", admin.User{}.GetMe)
		adAuth.Post("/category", admin.Category{}.Create)
	}
}
