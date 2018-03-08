package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler/article"
	"github.com/kataras/iris"
)

// ArticleRouter is
func ArticleRouter(party iris.Party) {
	o := party.Party("/article")
	{
		o.Get("/", article.Article{}.All)
		o.Get("/type/{id:string}", article.Article{}.AllType)
		o.Get("/{id:string}", article.Article{}.Get)
	}
	a := o.Party("/a")
	a.Use(jwt.JwtHandler.Serve)
	{
		a.Post("/", article.Article{}.Create)
		a.Delete("/", article.Article{}.Delete)
	}

}
