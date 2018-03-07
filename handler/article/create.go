package article

import (
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
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
