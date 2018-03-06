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

// Create is category new
func (c Category) Create(ctx iris.Context) {
	u := &models.Category{}
	if err := ctx.ReadJSON(u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.Name == "" {
		response.JSONError(ctx, "Name where?")
		return
	}

	cate, _ := models.Category{}.GetByName(u.Name)

	if cate.Name != "" {
		response.JSONError(ctx, "名称已存在")
		return
	}

	err := models.Category{}.Create(u)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, "创建成功")
}
