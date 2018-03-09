package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler/article"
	"anla.io/taizhou-fe-api/handler/comment"
	"anla.io/taizhou-fe-api/middleware"
	"github.com/kataras/iris"
)

// CommentRouter is
func CommentRouter(party iris.Party) {
	o := party.Party("/comment")
	{
		o.Get("/", article.Article{}.All)
		o.Get("/{id:string}", article.Article{}.Get)
	}
	a := o.Party("/a")
	a.Use(jwt.JwtHandler.Serve, middleware.User{}.GetUserCheck)
	{
		a.Post("/", comment.Comment{}.Save)
	}

}
