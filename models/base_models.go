package models

import (
	"time"
)

type (
	// IDAutoModel is
	IDAutoModel struct {
		ID uint `json:"id" sql:"AUTO_INCREMENT" gorm:"unique_index;not null;unique;primary_key;column:id"`
	}
	// IDModel is
	IDModel struct {
		ID string `json:"id" sql:"index"  gorm:"unique_index;not null;unique;primary_key;column:id"`
	}
	// UUIDModel is
	UUIDModel struct {
		UID string `json:"uid" sql:"index"  gorm:"unique_index;not null;unique;primary_key;column:uid"`
	}
)

type (
	// CreateModel is
	CreateModel struct {
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	}
	// UpdatedAtModel is
	UpdatedAtModel struct {
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	}
	// DeletedAtModel is
	DeletedAtModel struct {
		DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	}
	// TimeAllModel is
	TimeAllModel struct {
		CreateModel
		UpdatedAtModel
		DeletedAtModel
	}
	// UserIDModel is
	UserIDModel struct {
		UserID string `json:"user_id"`
	}
	// ParentIDModel is
	ParentIDModel struct {
		ParentID string `json:"parent_id"` //直接父评论的ID
	}
	// ArticleIDModel is
	ArticleIDModel struct {
		ArticleID string `json:"article_id"` //文章 ID
	}
	StatusModel struct {
		Status int `json:"status"`
	}
)

// BaseModel is
type BaseModel struct {
	IDModel
	TimeAllModel
}

// IDBaseModel is
type IDBaseModel struct {
	IDModel
	TimeAllModel
}

// UUIDBaseModel is
type UUIDBaseModel struct {
	UUIDModel
	TimeAllModel
}

// IDCreateModel is
type IDCreateModel struct {
	IDModel
	CreateModel
}

// PageModel is
type PageModel struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Size   int `json:"size"`
	Count  int `json:"count"`
	Num    int `json:"num"`
}

// QueryParams is
type QueryParams struct {
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	WareroomID int    `json:"wareroom_id"`
	Day        int    `json:"day"`
	ProductID  int    `json:"product_id"`
}

// QueryParamsTime is
type QueryParamsTime struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// PathParams is
type PathParams struct {
	ID uint64 `json:"id"`
}

type (
	TitleModel struct {
		Title string `json:"title"`
	}
)
