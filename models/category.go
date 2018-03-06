package models

import (
	"time"

	"anla.io/taizhou-fe-api/db"
	"github.com/houndgo/suuid"
	gm "github.com/jinzhu/gorm"
)

type (
	// Category is
	Category struct {
		BaseModel
		Name     string `json:"name" gorm:"type:varchar(100);unique"`
		Sort     int    `json:"sort"`
		Disabled bool   `json:"disabled" gorm:"default:'0'"`
	}
)

//BeforeSave is
func (a *Category) BeforeSave(scope *gm.Scope) (err error) {
	a.ID = suuid.New().String()
	return err
}

// Create is
func (a Category) Create(m *Category) error {
	var err error

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
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
func (a Category) GetAll() ([]Category, error) {
	var (
		data []Category
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}
