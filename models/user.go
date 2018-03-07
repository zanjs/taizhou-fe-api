package models

import (
	"time"

	"anla.io/taizhou-fe-api/db"
	"github.com/houndgo/suuid"
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
		Email        string `json:"email" gorm:"type:varchar(100);unique"`
		Password     string `json:"-"`
		Experience   uint   `json:"experience"`
		ArticleCount uint   `json:"article_count"` // 文章数
		CommentCount uint   `json:"comment_count"` // 评论数
		CollectCount uint   `json:"collect_count"` // 收藏数
		LaudCount    uint   `json:"laud_count"`    // 赞数
		AvatarURL    string `json:"avatar_url"`    //头像
		CoverURL     string `json:"cover_url"`     //个人主页背景图片URL
		Signature    string `json:"signature"`     //个人签名
		Role         int    `json:"role"`          //角色
		StatusModel
	}

	// UserShort is
	UserShort struct {
		UUIDModel
		UserName
	}
	// UserLogin is
	UserLogin struct {
		UserName
		Role     int    `json:"role"` //角色
		Email    string `json:"email" gorm:"type:varchar(100);unique"`
		Type     int    `json:"type" gorm:"-"` // 登陆类型 1 = 后台， 不填写为普通用户
		Password string `json:"password"`
	}
)

//TableName is set User's table name to be `users`
func (UserShort) TableName() string {
	return "users"
}

//BeforeSave is
// func (s *User) BeforeSave(scope *gm.Scope) (err error) {
// 	s.ID = suuid.New().String()
// 	return err
// }

// Create is user
func (s User) Create(m *User) error {
	var (
		err error
	)
	m.ID = suuid.New().String()
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

// GetByEmail is find user
func (s User) GetByEmail(data string) (User, error) {
	var (
		user User
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&user, "email = ?", data).Error; err != nil {
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

// GetAll is find
func (s User) GetAll(page *PageModel) ([]User, error) {
	var (
		data []User
		err  error
	)

	if page.Num < 1 {
		page.Num = 1
	}

	if page.Size == 0 {
		page.Size = 3
	}

	offset := (page.Num - 1) * page.Size

	tx := gorm.MysqlConn().Begin()

	if err = tx.Find(&data).Count(&page.Count).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	if err = tx.Order("created_at desc").Offset(offset).Limit(page.Size).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	tx.Commit()

	return data, err
}
