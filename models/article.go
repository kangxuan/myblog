package models

import "time"

type Article struct {
	ArticleId  int       `gorm:"primaryKey" json:"article_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Sort       int       `json:"sort"`
	IsDelete   int       `json:"is_delete"`
	CreateTime int       `json:"create_time"`
	UpdateTime int       `json:"update_time"`
	UpdateAt   time.Time `json:"update_at"`
}

type ArticleSaveColumns struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Sort       int    `json:"sort"`
	CategoryId int    `json:"category_id"`
	TagId      int    `json:"tag_id"`
}

type ArticleSearchColumns struct {
	ArticleId  int
	Title      string
	CategoryId int
	TagId      int
	Page       int
	PageSize   int
}

type ArticleListColumns struct {
	ArticleId     int    `json:"article_id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	CategoryId    int    `json:"category_id"`
	TagId         int    `json:"tag_id"`
	CreateTime    int    `json:"create_time"`
	CreateTimeStr string `json:"create_time_str"`
}
