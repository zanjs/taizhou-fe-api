package handler

import (
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"anla.io/taizhou-fe-api/utils"
	"github.com/kataras/iris"
)

type (
	// Register 注册接口
	Register struct {
		Controller
	}
)

func (re Register) mid(ctx iris.Context) (models.User, string) {
	u := &models.UserLogin{}
	user := models.User{}
	if err := ctx.ReadJSON(u); err != nil {
		response.JSONError(ctx, err.Error())
		return user, err.Error()
	}

	if u.Username == "" {
		return user, "Username where?"
	}
	if u.Email == "" {
		return user, "Email where?"
	}
	if u.Password == "" {
		return user, "Password where?"
	}
	user, _ = models.User{}.GetByUsername(u.Username)

	if user.ID != "" {
		return user, "用户名存在"
	}

	user, _ = models.User{}.GetByEmail(u.Email)

	if user.Email != "" {
		return user, "邮箱已存在"
	}

	user.Username = u.Username
	user.Role = u.Role
	user.Email = u.Email
	user.Password = utils.HashPassword(u.Password)
	return user, ""
}

// Add is add user
func (re Register) Add(ctx iris.Context) {

	user, errStr := re.mid(ctx)
	user.Role = 0

	if errStr != "" {
		response.JSONError(ctx, errStr)
		return
	}

	err := models.User{}.Create(&user)

	if err != nil {
		response.JSONError(ctx, "注册失败"+err.Error())
		return
	}

	response.JSON(ctx, user)
}

// AdminAdd is add user
func (re Register) AdminAdd(ctx iris.Context) {

	user, errStr := re.mid(ctx)

	if errStr != "" {
		response.JSONError(ctx, errStr)
		return
	}

	onUser := re.GetUser(ctx)

	if user.Role > onUser.Role {
		response.JSONError(ctx, "非法操作, 不得创建高于自己的角色")
		return
	}

	err := models.User{}.Create(&user)

	if err != nil {
		response.JSONError(ctx, "注册失败"+err.Error())
		return
	}

	response.JSON(ctx, user)
}
