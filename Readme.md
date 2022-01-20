
# go命令
> 1. 开启GOMODEL、设置GOPROXY
> 2. go mod init projectname   初始化项目
> 3. go get url  添加依赖
> 4. go mod tidy 检查&整理依赖

# 代码结构
go-gin-example/
├── conf
├── middleware
├── models
├── pkg
├── routers
└── runtime

* conf: 存储配置文件
* middleware: 应用中间件
* models: 应用数据库模型
* pkg: 第三方包
* routers: 路由逻辑处理
* runtime: 应用运行时数据

# go模块
> 1. import uri 引入网络依赖
> 2. replace uri => otheruri 替换指定依赖为本地目录/其他网络地址. 支持相对路径
> 3. 更换完依赖记得go mod tidy

# go配置加载包
> go get -u github.com/go-ini/ini