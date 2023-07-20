package dao

import (
	"errors"
	"myblog/models"
	"time"
)

type TagColumns struct {
	TagId   int
	TagName string
}

// GetTagList 获取
func GetTagList(params map[string]string) (tagList []*TagColumns, err error) {
	whereMaps := make(map[string]interface{})
	whereMaps["is_delete"] = 0
	if _, ok := params["tag_name"]; ok && params["tag_name"] != "" {
		whereMaps["tag_name"] = params["tag_name"]
	}
	result := models.Db.Table("tag").Where(whereMaps).Find(&tagList)
	if result.Error != nil {
		err = result.Error
		return
	}
	return tagList, nil
}

// CreateTag 创建标签
func CreateTag(tagParams *models.Tag) (id int, err error) {
	// 检查标签名是否存在
	var tagNameCount int64
	models.Db.Table("tag").Where("tag_name = ? and is_delete = 0", tagParams.TagName).Count(&tagNameCount)
	if tagNameCount > 0 {
		err = errors.New("标签名已存在")
		return
	}

	unix := time.Now().Unix()
	tag := models.Tag{
		TagName:    tagParams.TagName,
		CreateTime: int(unix),
		UpdateTime: int(unix),
	}
	result := models.Db.Table("tag").Create(&tag)
	if result.Error != nil {
		return 0, result.Error
	}

	return tag.TagId, nil
}

// UpdateTag 更新标签
func UpdateTag(id int, tagParams *models.Tag) error {
	// 检查id是否存在
	var tagIdCount, tagNameCount int64
	models.Db.Table("tag").Where("tag_id = ? and is_delete =0", id).Count(&tagIdCount)
	if tagIdCount == 0 {
		return errors.New("标签ID不存在")
	}
	models.Db.Table("tag").Where("tag_id <> ? and tag_name = ? and is_delete = 0", id, tagParams.TagName).Count(&tagNameCount)
	if tagNameCount > 0 {
		return errors.New("标签名已存在")
	}

	result := models.Db.Table("tag").Where("tag_id = ?", id).Updates(&models.Tag{
		TagName:    tagParams.TagName,
		UpdateTime: int(time.Now().Unix()),
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteTag 删除标签
func DeleteTag(id int) error {
	result := models.Db.Table("tag").Where("tag_id = ?", id).Updates(&models.Tag{
		IsDelete:   1,
		UpdateTime: int(time.Now().Unix()),
	})
	if result.Error != nil {
		return result.Error
	}

	// TODO：删除关联数据
	return nil
}
