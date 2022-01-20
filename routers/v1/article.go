package v1

import "github.com/gin-gonic/gin"

type Article struct {
	Model

	TagID int `json: "tag_id" gorm: "index"`
	Tag   Tag `json: "tag"`

	Title      string `json: "title"`
	Desc       string `json: "desc"`
	Content    string `json: "content"`
	CreateBy   string `json: "create_by"`
	ModifiedBy string `json: "modified_by"`
	State      int    `json: "state"`
}

func GetArticle(c *gin.Context) {

}

func GetArticles(c *gin.Context) {

}

func AddArticles(c *gin.Context) {

}

func EditArticles(c *gin.Context) {

}

func DeleteArticles(c *gin.Context) {

}
