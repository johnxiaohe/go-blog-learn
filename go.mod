module gin-example.com/v0

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.3
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

// replace (
// 	gin-example.com/v0/conf => ./conf
// 	gin-example.com/v0/middleware => ./middleware
// 	gin-example.com/v0/models => ./models
// 	gin-example.com/v0/pkg/setting => ./pkg/setting
// 	gin-example.com/v0/pkg/e => ./pkg/e
// 	gin-example.com/v0/pkg/util => ./pkg/util
// 	gin-example.com/v0/routers => ./routers
// 	gin-example.com/v0/runtime => ./runtime
// )
