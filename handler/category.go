package handler

import (
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

type (
	// Category is
	Category struct {
		Controller
	}
)

// GetAll is category new
func (c Category) GetAll(ctx iris.Context) {

	datas, err := models.Category{}.GetAll()
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, datas)
}
