package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	// gorm:index 表示该字段为索引.
	TagID int `json:"tag_id" gorm:"index"`
	// 嵌套Tag模型. 可自动关联查询
	Tag Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	return article.ID > 0
}

func ExistArticleByTitle(title string) bool {
	var article Article
	db.Select("title").Where("title = ?", title).First(&article)

	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// 预加载关联结果进行映射填充
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	// 关联查询
	db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	// 判断title是否是string类型(断言)
	title := data["title"].(string)

	if ExistArticleByTitle(title) {
		return false
	}
	db.Create(&Article{
		TagID:    data["tag_id"].(int),
		Title:    title,
		Desc:     data["desc"].(string),
		Content:  data["content"].(string),
		CreateBy: data["created_by"].(string),
		State:    data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}
