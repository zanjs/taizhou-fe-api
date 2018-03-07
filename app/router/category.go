package router

import (
	"anla.io/taizhou-fe-api/handler/category"
	"github.com/kataras/iris"
)

// CategoryRouter is
func CategoryRouter(party iris.Party) {
	o := party.Party("/category")
	{
		o.Get("/", category.Category{}.GetAll)
	}
}
