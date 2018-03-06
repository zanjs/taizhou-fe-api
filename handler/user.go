package handler

import (
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
