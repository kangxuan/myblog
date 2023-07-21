package models

import "time"

type TagRelation struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	TagId      int       `json:"tag_id"`
	Type       int       `json:"type"`
	RelationId int       `json:"relation_id"`
	IsDelete   int       `json:"is_delete"`
	CreateTime int       `json:"create_time"`
	UpdateTime int       `json:"update_time"`
	UpdateAt   time.Time `json:"update_at"`
}
