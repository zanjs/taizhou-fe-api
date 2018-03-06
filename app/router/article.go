package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler"
	"github.com/kataras/iris"
)

// ArticleRouter is
func ArticleRouter(party iris.Party) {
	o := party.Party("/article")
	{
		o.Get("/", handler.Article{}.All)
		o.Get("/{id:string}", handler.Article{}.Get)
	}
	a := o.Party("/a")
	a.Use(jwt.JwtHandler.Serve)
	{
		a.Post("/", handler.Article{}.Create)
	}

}
