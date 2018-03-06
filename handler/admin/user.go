package admin

import (
	"fmt"
	"strconv"
	"time"

	"anla.io/taizhou-fe-api/config"
	"anla.io/taizhou-fe-api/models"
	"anla.io/taizhou-fe-api/response"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
)

var jwtConfig = config.Config.JWT

// Login is login get token
func Login(ctx iris.Context) {

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

	user, _ := models.AdminUser{}.GetByUsername(u.Username)

	if user.ID == "" {
		response.JSONError(ctx, "用户名不存在")
		return
	}

	// hashPassword := utils.HashPassword(u.Password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err != nil {
		response.JSONError(ctx, "用户名或密码错误")
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	// claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtConfig.AdminSecret))
	if err != nil {
		response.JSONError(ctx, "err")
		return
	}

	response.JSON(ctx, t)
}

// User is
type User struct {
	Controller
}

// GetMe is
func (ctl User) GetMe(ctx iris.Context) {
	user := ctl.GetUser(ctx)

	datas, err := models.AdminUser{}.GetByID(user.ID)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, datas)
}

// GetAll is
func (ctl User) GetAll(ctx iris.Context) {
	pageNoStr := ctx.Request().FormValue("page_no")
	var pageNo int
	var err error
	if pageNo, err = strconv.Atoi(pageNoStr); err != nil {
		pageNo = 1
	}

	page := models.PageModel{}

	page.Num = pageNo

	datas, err := models.AdminUser{}.GetAll(&page)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	fmt.Println(datas)
	response.JSONPage(ctx, datas, page)
}
