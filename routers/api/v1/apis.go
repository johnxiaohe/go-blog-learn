package v1

import "github.com/gin-gonic/gin"

func InitApiV1(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	InitTagApis(apiV1)
	InitArticleApis(apiV1)
}
