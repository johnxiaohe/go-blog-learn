module gin-example.com/v0

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.3
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
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
