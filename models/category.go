package models

import (
	"time"

	"anla.io/taizhou-fe-api/db"
	"github.com/houndgo/suuid"
)

type (
	// Category is
	Category struct {
		BaseModel
		Name     string `json:"name" gorm:"type:varchar(100);unique"`
		Sort     int    `json:"sort"`
		Disabled bool   `json:"disabled" gorm:"default:'0'"`
		Article  []*Article
	}
	// CategoryArticle is
	CategoryArticle struct {
		CategoryID string `json:"category_id"`
		ArticleID  string `json:"article_id"`
	}
)

// Create is
func (a Category) Create(m *Category) error {
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

// Update is
func (a *Category) Update(m *Category) error {
	var err error

	a.UpdatedAt = time.Now()
	a.Name = m.Name
	a.Sort = m.Sort
	a.Disabled = m.Disabled

	tx := gorm.MysqlConn().Begin()

	if err = tx.Save(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// GetByID is find Category
func (a Category) GetByID(id string) (Category, error) {
	var (
		data Category
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&data, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// GetByName is
func (a Category) GetByName(name string) (Category, error) {
	var (
		data Category
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&data, "name = ?", name).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// GetAll is find
func (a Category) GetAll(page *PageModel) ([]Category, error) {
	var (
		data []Category
		err  error
	)

	tx := gorm.MysqlConn().Begin()

	if err = tx.Find(&data).Count(&page.Count).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	if err = tx.Order("sort desc,created_at desc").Offset(page.Offset).Limit(page.Size).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	tx.Commit()

	return data, err
}
