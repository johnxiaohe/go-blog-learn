package routers

import (
	"gin-example.com/v0/pkg/setting"
	v1 "gin-example.com/v0/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	
	r.GET("/AUTH", v1.GetAuth)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    "200",
			"data":    "null",
		})
	})

	v1.InitApiV1(r)

	return r
}
