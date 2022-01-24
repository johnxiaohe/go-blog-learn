package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// 定义加载配置的对象 和 变量
var (
	// 配置对象, 存储地址指针. 首字母大写表示公共属性
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse '/conf/app.ini': %+v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()

}

func LoadBase() {
	// 加载没有标签的 RUN_MODE. 并设置默认值为debug string
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	// 加载server标签的配置
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server configs' : %+v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app' : %+v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("abcdefg")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
