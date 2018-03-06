package models

import (
	"time"

	"anla.io/taizhou-fe-api/db"
	"github.com/houndgo/suuid"
	gm "github.com/jinzhu/gorm"
)

type (
	// AdminUser is
	AdminUser struct {
		BaseModel
		UserName
		Password string `json:"-"`
		StatusModel
		AvatarURL string `json:"avatar_url"` //头像
		CoverURL  string `json:"cover_url"`  //个人主页背景图片URL
	}
)

//BeforeSave is
func (s *AdminUser) BeforeSave(scope *gm.Scope) (err error) {
	s.ID = suuid.New().String()
	return err
}

// Create is user
func (s AdminUser) Create(m *AdminUser) error {
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
func (s AdminUser) GetByUsername(username string) (AdminUser, error) {
	var (
		user AdminUser
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
func (s AdminUser) GetByID(id string) (AdminUser, error) {
	var (
		user AdminUser
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
func (s AdminUser) GetAll(page *PageModel) ([]AdminUser, error) {
	var (
		data []AdminUser
		err  error
	)

	if page.Num < 1 {
		page.Num = 1
	}

	pageSize := 2
	offset := (page.Num - 1) * pageSize

	tx := gorm.MysqlConn().Begin()

	if err = tx.Find(&data).Count(&page.Count).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	if err = tx.Offset(offset).Limit(pageSize).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	tx.Commit()

	return data, err
}
