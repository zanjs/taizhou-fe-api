package admin

import (
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

	user, _ := models.User{}.GetByUsername(u.Username)

	if user.ID == "" {
		response.JSONError(ctx, "用户名不存在")
		return
	}

	if user.Role < models.UserRols.Edit {
		response.JSONError(ctx, "权限不够，请联系管理员")
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
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		response.JSONError(ctx, "err")
		return
	}

	response.JSON(ctx, t)
}
