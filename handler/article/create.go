package article

import (
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// Create is
func (ctl Article) Create(ctx iris.Context) {
	u, errStr := ctl.mid(ctx)
	if errStr != "" {
		response.JSONError(ctx, errStr)
		return
	}

	user := ctl.GetUser(ctx)

	u.UserID = user.ID

	err := models.Article{}.Create(&u)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSONSuccess(ctx)
}
