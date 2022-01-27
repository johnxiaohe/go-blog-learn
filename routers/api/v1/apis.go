package v1

import (
	"gin-example.com/v0/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitApiV1(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	InitTagApis(apiV1)
	InitArticleApis(apiV1)
}
