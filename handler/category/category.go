package category

import (
	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/middleware"
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

type (
	// Category is
	Category struct {
		handler.Controller
	}
)

// GetAll is category new
func (c Category) GetAll(ctx iris.Context) {

	page := middleware.GetPage(ctx)

	datas, err := models.Category{}.GetAll(&page)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSONPage(ctx, datas, page)
}
