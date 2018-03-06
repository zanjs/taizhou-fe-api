package models

import (
	"anla.io/taizhou-fe-api/db"
)

//CreateTable is init db table
func CreateTable() error {
	gorm.MysqlConn().AutoMigrate(&User{},
		&AppInfo{},
		&Article{},
		&ArticlePic{},
		&Category{},
		&AdminUser{},
		&Comment{})
	return nil
}
