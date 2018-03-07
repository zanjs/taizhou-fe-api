package admin

import (
	"anla.io/taizhou-fe-api/handler"
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

// Update is category update
func (c Category) Update(ctx iris.Context) {
	u := models.Category{}
	if err := ctx.ReadJSON(&u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.Name == "" {
		response.JSONError(ctx, "Name where?")
		return
	}

	if u.ID == "" {
		response.JSONError(ctx, "ID where?")
		return
	}

	cate, _ := models.Category{}.GetByID(u.ID)

	if cate.Name == "" {
		response.JSONError(ctx, "不存在")
		return
	}

	cateName, _ := models.Category{}.GetByName(u.Name)

	if cateName.Name != "" && u.ID != cateName.ID {
		response.JSONError(ctx, "名称已存在")
		return
	}

	err := cate.Update(&u)
	if err != nil {
		response.JSONError(ctx, err.Error()+"s")
		return
	}

	response.JSON(ctx, "更新成功")
}
