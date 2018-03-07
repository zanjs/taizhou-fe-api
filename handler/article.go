package handler

import (
	"fmt"
	"strconv"

	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

type (
	// Article is
	Article struct {
		Controller
	}
)

// Create is
func (ctl Article) Create(ctx iris.Context) {
	u := models.Article{}
	if err := ctx.ReadJSON(&u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.Content == "" {
		response.JSONError(ctx, "Content where?")
		return
	}

	if u.ContentType == 0 {
		response.JSONError(ctx, "ArticleType where?")
		return
	}

	user := ctl.GetUser(ctx)

	u.UserID = user.ID

	err := models.Article{}.Create(&u)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, u)
}

// All is
func (ctl Article) All(ctx iris.Context) {
	pageNoStr := ctx.Request().FormValue("page_no")
	var pageNo int
	var err error
	if pageNo, err = strconv.Atoi(pageNoStr); err != nil {
		pageNo = 1
	}

	page := models.PageModel{}

	page.Num = pageNo

	datas, err := models.Article{}.GetAll(&page)
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
