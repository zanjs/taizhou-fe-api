package handler

import (
	"fmt"
	"strconv"

	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// User is
type User struct {
	Controller
}

// GetMe is
func (ctl User) GetMe(ctx iris.Context) {
	user := ctl.GetUser(ctx)

	datas, err := models.User{}.GetByID(user.ID)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, datas)
}

// GetAll is
func (ctl User) GetAll(ctx iris.Context) {
	pageNoStr := ctx.Request().FormValue("page_no")
	var pageNo int
	var err error
	if pageNo, err = strconv.Atoi(pageNoStr); err != nil {
		pageNo = 1
	}

	page := models.PageModel{}

	page.Num = pageNo

	datas, err := models.User{}.GetAll(&page)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	fmt.Println(datas)
	response.JSONPage(ctx, datas, page)
}
