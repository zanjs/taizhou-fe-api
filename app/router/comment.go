package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/handler/comment"
	"github.com/kataras/iris"
)

// CommentRouter is
func CommentRouter(party iris.Party) {
	o := party.Party("/comment")
	{
		o.Get("/", handler.Article{}.All)
		o.Get("/{id:string}", handler.Article{}.Get)
	}
	a := o.Party("/a")
	a.Use(jwt.JwtHandler.Serve)
	{
		a.Post("/", comment.Comment{}.Save)
	}

}
