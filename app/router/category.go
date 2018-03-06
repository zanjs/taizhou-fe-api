package router

import (
	"anla.io/taizhou-fe-api/handler"
	"github.com/kataras/iris"
)

// CategoryRouter is
func CategoryRouter(party iris.Party) {
	o := party.Party("/category")
	{
		o.Get("/", handler.Category{}.GetAll)
	}
}
