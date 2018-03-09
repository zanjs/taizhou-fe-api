package middleware

import (
	"errors"
	"fmt"

	"anla.io/taizhou-fe-api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// User is
type User struct{}

func (c User) claimsOK(ctx iris.Context) (models.User, error) {
	user := models.User{}
	userJwt := ctx.Values().Get("jwt").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	// userID := uint(claims["id"].(float64))
	fmt.Println(claims)

	id := claims["id"]
	role := claims["role"]

	if id == "" || role == "" || id == nil || role == nil {
		fmt.Println("user 没有通过")
		return user, errors.New("user 没有通过")
	}
	user.ID = id.(string)
	user.Role = int(role.(float64))
	return user, nil
}

// GetUserCheck 验证用户信息
func (c User) GetUserCheck(ctx iris.Context) {

	_, err := c.claimsOK(ctx)
	if err != nil {
		return
	}
	ctx.Next()
}

// GetUser 获取用户信息
func (c User) GetUser(ctx iris.Context) models.User {
	user, _ := c.claimsOK(ctx)
	return user
}
