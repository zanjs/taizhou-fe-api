package models

import (
	"fmt"
	"time"

	"anla.io/taizhou-fe-api/db"
	"github.com/houndgo/suuid"
	"github.com/theplant/batchputs"
)

type (
	// ArticleType is
	articleType struct {
		New    int
		Bubble int
	}
)

// ArticleType 发帖类型
var ArticleType = articleType{
	Bubble: 1, // 冒泡， 用户发帖
	New:    2, // 专家发帖
}

type (

	// Article is
	Article struct {
		IDModel
		TimeAllModel
		User          User          `gorm:"Table:user;foreignkey:UserID;AssociationForeignKey:id" json:"user,omitempty"`
		Pics          []*ArticlePic `json:"pics,omitempty"`
		UserID        string        `json:"user_id"`
		ContentType   uint          `json:"content_type"`
		Title         string        `json:"title" gorm:"type:varchar(100)"`
		Content       string        `json:"content" gorm:"type:text"`
		ViewCount     int           `json:"view_count"`
		CommentCount  uint          `json:"comment_count"`
		Categories    []Category    `gorm:"many2many:article_category;ForeignKey:ID;AssociationForeignKey:ID" json:"categories"`
		LastUserID    uint          `json:"last_user_id"` //最后一个回复话题的人
		LastUser      User          `json:"last_user"`
		LastCommentAt *time.Time    `json:"last_comment_at"`
		Disabled      bool          `json:"disabled" gorm:"default:'0'"`
		Comments      []*Comment    `json:"comments" gorm:"-"`
	}
)

//BeforeSave is
// func (a *Article) BeforeSave(scope *gm.Scope) (err error) {
// 	a.ID = suuid.New().String()
// 	return err
// }

// Create is
func (a Article) Create(m *Article) error {
	var err error
	m.ID = suuid.New().String()
	rows := [][]interface{}{}

	// pics := &m.Pics

	for i := 0; i < len(m.Pics); i++ {
		m.Pics[i].ArticleID = m.ID
		m.Pics[i].CreatedAt = time.Now()
		uid := suuid.New().String()
		m.Pics[i].ID = uid
		rows = append(rows, []interface{}{
			m.ID,
			m.Pics[i].CreatedAt,
			m.Pics[i].Src,
			uid,
		})
	}

	columns := []string{"article_id", "created_at", "src", "id"}
	dialect := "mysql"

	err = batchputs.Put(gorm.MysqlConn().DB(), dialect, "article_pics", "article_id", columns, rows)
	if err != nil {
		panic(err)
	}

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err
}

// GetAll is find
func (a Article) GetAll(page *PageModel) ([]Article, error) {
	var (
		data []Article
		err  error
	)

	if page.Num < 1 {
		page.Num = 1
	}

	pageSize := 2
	offset := (page.Num - 1) * pageSize

	tx := gorm.MysqlConn().Begin()

	if err = tx.Preload("User").Preload("Pics").Find(&data).Count(&page.Count).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	if err = tx.Offset(offset).Limit(pageSize).Preload("User").Preload("Pics").Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	// tx.Model(&data).Related(&pics)
	// tx.Model(&data).Related(&pics, "Pics")
	// tx.Model(&data).Association("Pics")
	tx.Commit()

	return data, err
}

// GetAllNews is find
func (a Article) GetAllNews(page *PageModel) ([]Article, error) {
	var (
		data []Article
		err  error
	)

	if page.Num < 1 {
		page.Num = 1
	}

	pageSize := 2
	offset := (page.Num - 1) * pageSize

	tx := gorm.MysqlConn().Begin()

	if err = tx.Preload("User").Preload("Pics").Where("content_type = ?", ArticleType.New).Find(&data).Count(&page.Count).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	if err = tx.Offset(offset).Limit(pageSize).Preload("User").Preload("Pics").Where("content_type = ?", ArticleType.New).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	// tx.Model(&data).Related(&pics)
	// tx.Model(&data).Related(&pics, "Pics")
	// tx.Model(&data).Association("Pics")
	tx.Commit()

	return data, err
}

// Get is find
func (a Article) Get(id string) (Article, error) {
	var (
		data Article
		err  error
	)

	tx := gorm.MysqlConn().Begin()

	if err = tx.Where("id = ?", id).First(&data).Error; err != nil {
		fmt.Println(data)
		return data, err
	}
	tx.Commit()

	return data, err
}

// Update 更新保存
func (a Article) Update(data *Article) error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err
}
