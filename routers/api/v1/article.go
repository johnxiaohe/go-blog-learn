package v1

import (
	"github.com/gin-gonic/gin"
)

func getArticle(c *gin.Context) {

}

func getArticles(c *gin.Context) {

}

func addArticles(c *gin.Context) {

}

func editArticles(c *gin.Context) {

}

func deleteArticles(c *gin.Context) {

}

func InitArticleApis(apiV1 *gin.RouterGroup) {
	//获取文章列表
	apiV1.GET("/articles", getArticles)
	//获取指定文章
	apiV1.GET("/article/:id", getArticle)
	//新建文章
	apiV1.POST("/article", addArticles)
	//更新指定文章
	apiV1.PUT("/article/:id", editArticles)
	//删除指定文章
	apiV1.DELETE("/article/:id", deleteArticles)
}
