package middleware

import (
	"strconv"

	"anla.io/taizhou-fe-api/models"
	"github.com/kataras/iris"
)

// GetPage 获取分页数
func GetPage(ctx iris.Context) models.PageModel {
	pageNoStr := ctx.Request().FormValue("page_no")
	var pageNo int
	var err error
	if pageNo, err = strconv.Atoi(pageNoStr); err != nil {
		pageNo = 1
	}

	page := models.PageModel{}

	page.Num = pageNo

	if page.Num < 1 {
		page.Num = 1
	}

	if page.Size == 0 {
		page.Size = 3
	}

	page.Offset = (page.Num - 1) * page.Size

	return page
}
