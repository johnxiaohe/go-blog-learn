package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name" valid:"Required; MaxSize(100)"`
	CreatedBy  string `json:"createdBy" valid:"Required; MaxSize(100)"`
	ModifiedBy string `json:"modifiedBy"`
	State      int    `json:"state,string" valid:"Required; Range(0,1)"`
}

// 声明Tag的方法
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 查询标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)

	return tag.ID > 0
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}

func AddTag(tag *Tag) bool {
	if !ExistTagByName(tag.Name) {
		db.Create(&Tag{
			Name:      tag.Name,
			State:     tag.State,
			CreatedBy: tag.CreatedBy,
		})
	}

	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true
}
