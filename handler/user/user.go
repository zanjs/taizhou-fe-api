package user

import (
	"fmt"

	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/middleware"
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// User is
type User struct {
	handler.Controller
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
	page := middleware.GetPage(ctx)

	datas, err := models.User{}.GetAll(&page)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	fmt.Println(datas)
	response.JSONPage(ctx, datas, page)
}
