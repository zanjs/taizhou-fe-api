package article

import (
	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

type (
	// Article is
	Article struct {
		handler.Controller
	}
)

func (ctl Article) mid(ctx iris.Context) (models.Article, string) {
	u := models.Article{}
	if err := ctx.ReadJSON(&u); err != nil {
		return u, err.Error()
	}

	if u.Content == "" {
		return u, "Content where?"
	}

	if u.ContentType == 0 {
		return u, "ArticleType where?"
	}

	return u, ""
}

// Update is
func (ctl Article) Update(ctx iris.Context) {
	u, errStr := ctl.mid(ctx)
	if errStr != "" {
		response.JSONError(ctx, errStr)
		return
	}

	err := models.Article{}.Create(&u)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, "发布成功")
}
