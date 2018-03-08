package user

import (
	"fmt"

	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/middleware"
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"anla.io/taizhou-fe-api/utils"
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

// Update is
func (ctl User) Update(ctx iris.Context) {

	u := models.UserLogin{}
	if err := ctx.ReadJSON(&u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	fmt.Println("uuuuuuuuuuuuuuuuuuuuu")
	fmt.Println(u)

	if u.Password != "" {
		u.Password = utils.HashPassword(u.Password)
	}

	userJwt := ctl.GetUser(ctx)

	datas, err := models.User{}.GetByID(userJwt.ID)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	err = datas.Update(&u)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, "更新成功")
}
