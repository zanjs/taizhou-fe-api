package models

import (
	"time"

	"anla.io/taizhou-fe-api/db"
	"github.com/houndgo/suuid"
	gm "github.com/jinzhu/gorm"
)

type (
	// UserName is
	UserName struct {
		Username string `json:"username" gorm:"type:varchar(100);unique"`
	}
	// User is
	User struct {
		BaseModel
		UserName
		Email        string `json:"-" gorm:"type:varchar(100);unique"`
		Password     string `json:"-"`
		Experience   uint   `json:"experience"`
		ArticleCount uint   `json:"articleCount"` // 文章数
		CommentCount uint   `json:"commentCount"` // 评论数
		CollectCount uint   `json:"collectCount"` // 收藏数
		LaudCount    uint   `json:"laudCount"`    // 赞数
		AvatarURL    string `json:"avatarURL"`    //头像
		CoverURL     string `json:"coverURL"`     //个人主页背景图片URL
		Signature    string `json:"signature"`    //个人签名
	}

	// UserShort is
	UserShort struct {
		UUIDModel
		UserName
	}
	// UserLogin is
	UserLogin struct {
		UserName
		Password string `json:"password"`
	}
)

//TableName is set User's table name to be `users`
func (UserShort) TableName() string {
	return "users"
}

//BeforeSave is
func (s *User) BeforeSave(scope *gm.Scope) (err error) {
	s.ID = suuid.New().String()
	return err
}

// Create is user
func (s User) Create(m *User) error {
	var (
		err error
	)
	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// GetByUsername is find user
func (s User) GetByUsername(username string) (User, error) {
	var (
		user User
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&user, "username = ?", username).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}

// GetByID is find user
func (s User) GetByID(id string) (User, error) {
	var (
		user User
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&user, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}
