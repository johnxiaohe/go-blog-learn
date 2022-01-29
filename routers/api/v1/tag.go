package v1

import (
	"log"
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
	var tag models.Tag
	c.BindJSON(&tag)
	log.Printf("%+v\n", &tag)

	valid := validation.Validation{}
	ok, _ := valid.Valid(&tag)

	code := e.INVALID_PARAMS

	if ok {
		if !models.ExistTagByName(tag.Name) {
			code = e.SUCCESS
			models.AddTag(&tag)
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
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

//删除文章标签
func deleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func InitTagApis(apiv1 *gin.RouterGroup) {
	//获取标签列表
	apiv1.GET("/tags", getTags)
	//新建标签
	apiv1.POST("/tag", addTag)
	//更新指定标签
	apiv1.PUT("/tag/:id", editTag)
	//删除指定标签
	apiv1.DELETE("/tag/:id", deleteTag)
}
