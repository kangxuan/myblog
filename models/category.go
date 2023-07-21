package models

import "time"

type Category struct {
	CategoryId   int       `gorm:"primaryKey" json:"category_id"`
	CategoryType int       `json:"category_type"`
	CategoryName string    `json:"category_name"`
	ParentId     int       `json:"parent_id"`
	IsDelete     int       `json:"is_delete"`
	CreateTime   int       `json:"create_time"`
	UpdateTime   int       `json:"update_time"`
	UpdateAt     time.Time `json:"update_at"`
}

type CategorySearchColumns struct {
	CategoryName string
	Page         int
	PageSize     int
}
