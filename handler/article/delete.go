package article

import (
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

// Delete is
func (ctl Article) Delete(ctx iris.Context) {
	u := models.Article{}
	if err := ctx.ReadJSON(&u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.ID == "" {
		response.JSONError(ctx, "非法提交")
		return
	}

	user := ctl.GetUser(ctx)

	if user.ID != u.UserID || user.Role < models.UserRols.Edit {
		response.JSONError(ctx, "没有权限")
		return
	}

	datas, err := models.Article{}.Get(u.ID)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	if datas.ID == "" {
		response.JSONError(ctx, "不存在")
		return
	}

	err = datas.Delete()
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, user)
}
