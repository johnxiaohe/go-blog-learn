package v1

import (
	"net/http"

	"gin-example.com/v0/models"
	"gin-example.com/v0/pkg/e"
	"gin-example.com/v0/pkg/setting"
	"gin-example.com/v0/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取多个文章标签
func getTags(c *gin.Context) {
	name := c.Query("name")

	// 查询条件对象
	maps := make(map[string]interface{})
	// 查询结果对象
	data := make(map[string]interface{})

	// 查看是否有查询名称字段
	if name != "" {
		maps["name"] = name
	}

	// 查询状态字段
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//新增文章标签
func addTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符串")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

//修改文章标签
func editTag(c *gin.Context) {

}

//删除文章标签
func deleteTag(c *gin.Context) {

}

func InitTagApis(apiv1 *gin.RouterGroup) {
	//获取标签列表
	apiv1.GET("/tags", getTags)
	//新建标签
	apiv1.POST("/tags", addTag)
	//更新指定标签
	apiv1.PUT("/tags/:id", editTag)
	//删除指定标签
	apiv1.DELETE("/tags/:id", deleteTag)
}
