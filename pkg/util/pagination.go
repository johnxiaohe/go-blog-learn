package util

import (
	"gin-example.com/v0/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取当前页偏移量
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
