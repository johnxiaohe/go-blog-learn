package routers

import (
	"gin-example.com/v0/pkg/setting"
	v1 "gin-example.com/v0/routers/V1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    "200",
			"data":    "null",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/article/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/article", v1.AddArticles)
		//更新指定文章
		apiV1.PUT("/article/:id", v1.EditArticles)
		//删除指定文章
		apiV1.DELETE("/article/:id", v1.DeleteArticles)
	}

	return r
}
