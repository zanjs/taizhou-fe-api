package app

import (
	"anla.io/taizhou-fe-api/app/router"
	"anla.io/taizhou-fe-api/config"
	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var (
	appConf = config.Config.APP
)

// InitApp is
func InitApp() {
	app := iris.New()
	app.Use(crs)
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.Before)

	// attach the file as logger, remember, iris' app logger is just an io.Writer.
	app.Logger().SetOutput(config.NewLogFile())
	// app.Use(iris.Gzip)
	app.StaticWeb("/uploads", "./uploads")
	app.Options("/*", handler.OptionsHandler)

	app.Get("/", handler.IndexHand)
	app.Get("/init", handler.InitTable)
	app.Get("/suuid", handler.UUID{}.Create)

	v1 := app.Party("/api/v1")
	router.AdminRouter(v1)
	router.UserRouter(v1)
	router.ArticleRouter(v1)
	router.CategoryRouter(v1)
	router.CommentRouter(v1)
	router.FileRouter(v1)

	// navigate to defafult config http://localhost:8080
	if err := app.Run(iris.Addr(":"+appConf.Port), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}
