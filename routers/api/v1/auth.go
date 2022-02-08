package v1

import (
	"net/http"

	"gin-example.com/v0/models"
	"gin-example.com/v0/pkg/e"
	"gin-example.com/v0/pkg/logging"
	"gin-example.com/v0/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var json Auth
	c.BindJSON(&json)
	// log.Printf("%v\n", &json)
	valid := validation.Validation{}
	ok, _ := valid.Valid(&json)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(json.Username, json.Password)
		if isExist {
			token, err := util.GenerateToken(json.Username, json.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
