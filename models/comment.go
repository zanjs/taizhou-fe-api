package models

import (
	"time"

	"anla.io/taizhou-fe-api/db"
	"github.com/houndgo/suuid"
)

type (
	// Comment is articel comment
	Comment struct {
		BaseModel
		UserIDModel
		ArticleIDModel
		Content string `json:"content"`
		ParentIDModel
		User User `json:"user"`
	}
)

const (
	// CommentVerifying 审核中
	CommentVerifying = 1

	// CommentVerifySuccess 审核通过
	CommentVerifySuccess = 2

	// CommentVerifyFail 审核未通过
	CommentVerifyFail = 3
)

// Create is
func (a Comment) Create(m *Comment) error {
	var err error
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
