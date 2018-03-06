package handler

import (
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"anla.io/taizhou-fe-api/utils"
	"github.com/kataras/iris"
)

type (
	// Register 注册接口
	Register struct{}
)

// Add is add user
func (re Register) Add(ctx iris.Context) {
	u := &models.UserLogin{}
	if err := ctx.ReadJSON(u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.Username == "" {
		response.JSONError(ctx, "Username where?")
		return
	}

	if u.Password == "" {
		response.JSONError(ctx, "Password where?")
		return
	}

	user, _ := models.User{}.GetByUsername(u.Username)

	if user.ID != "" {
		response.JSONError(ctx, "用户名存在")
		return
	}

	user.Username = u.Username
	user.Password = utils.HashPassword(u.Password)

	err := models.User{}.Create(&user)

	if err != nil {
		response.JSONError(ctx, "注册失败")
	}

	response.JSON(ctx, user)
}
