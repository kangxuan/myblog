package models

import "time"

type CategoryRelation struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	CategoryId int       `json:"category_id"`
	Type       int       `json:"type"`
	RelateId   int       `json:"relate_id"`
	IsDelete   int       `json:"is_delete"`
	CreateTime int       `json:"create_time"`
	UpdateTime int       `json:"update_time"`
	UpdateAt   time.Time `json:"update_at"`
}
