package models

import "time"

type Category struct {
	CategoryId   int       `db:"category_id" json:"category_id"`
	CategoryType int       `db:"category_type" json:"category_type"`
	CategoryName string    `db:"category_name" json:"category_name"`
	ParentId     int       `db:"parent_id" json:"parent_id"`
	IsDelete     int       `db:"is_delete" json:"is_delete"`
	CreateTime   int       `db:"create_time" json:"create_time"`
	UpdateTime   int       `db:"update_time" json:"update_time"`
	UpdateAt     time.Time `db:"update_at" json:"update_at"`
}
