package comment

import (
	"fmt"

	"anla.io/taizhou-fe-api/handler"
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	"github.com/kataras/iris"
)

type (
	// Comment is
	Comment struct {
		handler.Controller
	}
)

// Save is
func (ctl Comment) Save(ctx iris.Context) {
	u := models.Comment{}
	if err := ctx.ReadJSON(&u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	fmt.Println(u.Content)

	if u.Content == "" {
		response.JSONError(ctx, "Content where?")
		return
	}

	if u.ArticleID == "" {
		response.JSONError(ctx, "ArticleID where?")
		return
	}

	articel, err := models.Article{}.Get(u.ArticleID)
	if err != nil {
		response.JSONError(ctx, "ArticleID no find")
		return
	}

	user := ctl.GetUser(ctx)

	if user.ID == "" {
		response.JSONError(ctx, "jwt no find")
	}

	fmt.Println(user)

	u.UserID = user.ID

	err = models.Comment{}.Create(&u)
	if err != nil {
		response.JSONError(ctx, "留言失败"+err.Error())
		return
	}

	articel.CommentCount++

	err = articel.Update(&articel)

	if err != nil {
		fmt.Println("更新留言数量错误:" + err.Error())
	}

	response.JSON(ctx, "发表成功，能力加1")
}
