package models

import (
	"fmt"
	"log"

	// gorm隐式依赖mysql驱动依赖
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin-example.com/v0/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {

	// 预声明变量
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	// 获取数据库配置片段
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %+v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// 开启db连接. Sprintf为格式化输出字符串
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))

	if err != nil {
		log.Println(err)
	}

	// 获取表名时的插件方法
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// 设置连接数等属性
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	// defer 关键字: 在当前函数返回前执行指定的函数
	defer db.Close()
}
