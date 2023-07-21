package dao

import (
	"errors"
	"gorm.io/gorm"
	"myblog/models"
	"myblog/pkg/constant"
	"myblog/pkg/util"
	"time"
)

// GetArticleList 获取文章列表
func GetArticleList(searchParams *models.ArticleSearchColumns) (articleList []*models.ArticleListColumns, err error) {
	query := getArticleQuery(searchParams)
	if err = util.Pagination(query, searchParams.Page, searchParams.PageSize).Scan(&articleList).Error; err != nil {
		return
	}
	for k, v := range articleList {
		articleList[k].CreateTimeStr = util.TimeToString(v.CreateTime)
		articleList[k].Content = util.Utf8Substring(articleList[k].Content, 0, 50) + "[...]"
	}
	return
}

// GetArticlePage 获取文章分页
func GetArticlePage(searchParams *models.ArticleSearchColumns) (page map[string]interface{}) {
	query := getArticleQuery(searchParams)
	return util.Page(query, searchParams.Page, searchParams.PageSize)
}

// getArticleQuery 获取文章查询器
func getArticleQuery(searchParams *models.ArticleSearchColumns) *gorm.DB {
	query := models.Db.Table("article").Joins("left join category_relation on article.article_id = category_relation.relate_id and category_relation.type = ? and category_relation.is_delete = ?", constant.CategoryRelationTypeArticle, 0)
	query.Joins("left join tag_relation on article.article_id = tag_relation.relate_id and tag_relation.type = ? and tag_relation.is_delete = ?", constant.TagRelationTypeArticle, 0)
	query.Select("article.article_id, article.title, article.content, category_relation.category_id, tag_relation.tag_id, article.create_time")
	if searchParams.Title != "" {
		query.Where("article.title like ?", "%"+searchParams.Title+"%")
	}
	if searchParams.TagId != 0 {
		query.Where("tag_relation.tag_id = ?", searchParams.TagId)
	}
	if searchParams.CategoryId != 0 {
		query.Where("category_relation.category_id = ?", searchParams.CategoryId)
	}
	if searchParams.ArticleId != 0 {
		query.Where("article.article_id = ?", searchParams.ArticleId)
	}
	return query
}

// GetArticle 获取文章详情
func GetArticle(id int) (article *models.ArticleListColumns, err error) {
	query := getArticleQuery(&models.ArticleSearchColumns{
		ArticleId: id,
	})
	result := query.Scan(&article)
	if err = result.Error; err != nil {
		return
	}
	if article.ArticleId != 0 {
		article.CreateTimeStr = util.TimeToString(article.CreateTime)
	}
	return
}

// CreateArticle 创建文章
func CreateArticle(articleParams *models.ArticleSaveColumns) (int, error) {
	currentTime := int(time.Now().Unix())
	// 添加文章
	createArticle := models.Article{
		Title:      articleParams.Title,
		Sort:       articleParams.Sort,
		Content:    articleParams.Content,
		CreateTime: currentTime,
		UpdateTime: currentTime,
	}
	err := models.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Table("article").Create(&createArticle)
		if result.Error != nil {
			return result.Error
		}
		// 添加文章分类关联
		createCategoryRelation := models.CategoryRelation{
			CategoryId: articleParams.CategoryId,
			Type:       constant.CategoryRelationTypeArticle,
			RelateId:   createArticle.ArticleId,
			CreateTime: currentTime,
			UpdateTime: currentTime,
		}
		if err := tx.Table("category_relation").Create(&createCategoryRelation).Error; err != nil {
			return err
		}
		// 添加标签关联
		if articleParams.TagId > 0 {
			createTagRelation := models.TagRelation{
				TagId:      articleParams.TagId,
				Type:       constant.TagRelationTypeArticle,
				RelationId: createArticle.ArticleId,
				CreateTime: currentTime,
				UpdateTime: currentTime,
			}
			if err := tx.Table("tag_relation").Create(&createTagRelation).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return 0, err
	}
	return createArticle.ArticleId, nil
}

// UpdateArticle 修改文章
func UpdateArticle(id int, articleParams *models.ArticleSaveColumns) error {
	// 判断文章ID是否存在
	var existedCount int64
	models.Db.Table("article").Where("article_id = ? and is_delete=0", id).Count(&existedCount)
	if existedCount == 0 {
		return errors.New("文章ID不存在")
	}
	currentTime := int(time.Now().Unix())
	err := models.Db.Transaction(func(tx *gorm.DB) error {
		// 先删除所有关联的记录
		models.Db.Table("category_relation").Where("relate_id = ? and type = ?", id, constant.CategoryRelationTypeArticle).Update("is_delete", 1)
		models.Db.Table("tag_relation").Where("relate_id = ? and type type = ?", id, constant.TagRelationTypeArticle).Update("is_delete", 1)
		// 修改文章
		updateArticle := models.Article{
			Title:      articleParams.Title,
			Sort:       articleParams.Sort,
			Content:    articleParams.Content,
			UpdateTime: currentTime,
		}
		result := tx.Table("article").Where("article_id = ?", id).Updates(&updateArticle)
		if result.Error != nil {
			return result.Error
		}
		// 修改文章分类关联
		createCategoryRelation := models.CategoryRelation{
			CategoryId: articleParams.CategoryId,
			Type:       constant.CategoryRelationTypeArticle,
			RelateId:   id,
			CreateTime: currentTime,
			UpdateTime: currentTime,
		}
		if err := tx.Table("category_relation").Create(&createCategoryRelation).Error; err != nil {
			return err
		}
		// 添加标签关联
		if articleParams.TagId > 0 {
			createTagRelation := models.TagRelation{
				TagId:      articleParams.TagId,
				Type:       constant.TagRelationTypeArticle,
				RelationId: id,
				CreateTime: currentTime,
				UpdateTime: currentTime,
			}
			if err := tx.Table("tag_relation").Create(&createTagRelation).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteArticle 删除文章
func DeleteArticle(id int) error {
	return models.Db.Transaction(func(tx *gorm.DB) error {
		currentTime := int(time.Now().Unix())
		// 删除文章
		if err := tx.Table("article").Where("article_id = ?", id).Updates(&models.Article{
			IsDelete:   1,
			UpdateTime: currentTime,
		}).Error; err != nil {
			return err
		}
		// 删除关联
		models.Db.Table("category_relation").Where("relate_id = ? and type = ?", id, constant.CategoryRelationTypeArticle).Update("is_delete", 1)
		models.Db.Table("tag_relation").Where("relate_id = ? and type type = ?", id, constant.TagRelationTypeArticle).Update("is_delete", 1)
		return nil
	})
}
