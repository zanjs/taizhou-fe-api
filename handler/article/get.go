package article

import (
	"fmt"
	"strconv"

	"anla.io/taizhou-fe-api/middleware"
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// All is
func (ctl Article) All(ctx iris.Context) {
	page := middleware.GetPage(ctx)

	datas, err := models.Article{}.GetAll(&page)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	fmt.Println(datas)
	response.JSONPage(ctx, datas, page)
}

// AllType is
func (ctl Article) AllType(ctx iris.Context) {
	page := middleware.GetPage(ctx)

	cType, _ := strconv.Atoi(ctx.Params().Get("id"))

	datas, err := models.Article{}.GetAllType(&page, cType)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	fmt.Println(datas)
	response.JSONPage(ctx, datas, page)
}

// Get is
func (ctl Article) Get(ctx iris.Context) {
	id := ctx.Params().Get("id")

	datas, err := models.Article{}.Get(id)
	if err != nil {
		response.JSONError(ctx, err.Error()+" id:"+id)
		return
	}

	user, err := models.User{}.GetByID(datas.UserID)
	if err != nil {
		fmt.Println("未找到关联的用户", err.Error())
	}

	datas.User = user

	response.JSON(ctx, datas)
}
