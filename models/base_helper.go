package models

import (
	"anla.io/taizhou-fe-api/db"
)

func helpCreate(m *struct{}) error {
	var (
		err error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
