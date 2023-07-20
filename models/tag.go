package models

import "time"

type Tag struct {
	TagId      int       `json:"tag_id"`
	TagName    string    `json:"tag_name"`
	IsDelete   int       `json:"is_delete"`
	CreateTime int       `json:"create_time"`
	UpdateTime int       `json:"update_time"`
	UpdateAt   time.Time `json:"update_at"`
}
