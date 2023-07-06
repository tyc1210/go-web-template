## go web开发模板项目
### 介绍
该项目适用于初学者使用 go 语言进行 web 开发
实现功能有：用户注册、登录、查询
技术：gin、swagger、gorm

### 项目目录介绍

- configs 配置文件 yaml
- global 声明全局变量
- internal
  - dao 操作数据库
  - middleware 中间件 例如：权限验证、跨域等
  - model 数据库对应实体类
  - routers
    - handler 各种处理器可类比 controller 层(参数绑定、调用service) 
    - router gin初始化，项目路由相关
  - service
  - setting 配置文件对应实体类、配置viper
- pkg
  - app
    - common_result 全局返回
  - errcode 自定义返回码
  - logger 日志
  - util 工具类
- storage 存储上传文件、日志等

### 数据库
> go get github.com/jinzhu/gorm
> github.com/go-sql-driver/mysql

- 定义方法从配置读取数据库配置信息返回 *gorm.DB
- global中 声明全局变量类型为 *gorm.DB
- 在 main 方法中初始化 global 数据库变量

### 业务处理流程
以用户管理举例

- handler中编写 User 结构体，为其添加增删改查方法，方法参数都包含 *gin.Context
- 从 *gin.Context 获取参数信息调用 service 
- 声明 Service 结构体
  - 属性包含：context.Context、 *dao.Dao
  - 提供 New 方法初始化属性
  - 为 Service 添加各种供 Controller 曾调用的方法
- 声明 Dao 结构体
  - 属性包含 *gorm.db
  - 提供 New 方法，初始化属性
  - 提供各种供 Service 层调用的方法


### 配置文件
> go get github.com/spf13/viper 

- 项目下创建configs文件夹，编写 yaml 配置文件
- 在 internal 文件夹下创建 setting 文件夹编写 viper 配置文件读取信息以及对应的实体类信息
- 在 global 文件夹下声明指向配置对应实体类的指针类型的全局变量
- main 方法中对配置viper开启读取环境变量，对配置相关全局变量进行初始化
开启读取环境变量后，若环境变量中配置了 SERVER_HTTPPORT 则会替换掉配置文件中的Server.HttpPort



### swagger
```shell
$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
$ go get -u github.com/swaggo/gin-swagger@v1.2.0 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template
```
到$gopath/pkg/mod/github.com/swaggo/swag@v1.6.5/cmd/swag下运行go build 生成swag.exe 放到$gopath/bin目录下

swagger常用注解

- @Summary 摘要
- @Produce API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
- @Param 参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
- @Success 响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
- @Failure 响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
- @Router 路由，从左到右分别为：路由地址，HTTP 方法

使用步骤

- 在处理器方法上编写注解
- 执行 swag init (在/docs文件夹下生成swag相关文件)
- 在 router.go 中进行默认初始化和注册对应的 swag 路由
- 注意一定要在 router.go 文件中导入 docs.go 例如： _ "go-web-template/docs"
- 启动服务 访问：http://127.0.0.1:{port}/swagger/index.html

### jwt
> go get github.com/dgrijalva/jwt-go

- 编写 jwt 工具类用来生成、解析token
- 编写 jwt 中间件，从请求头中获取token并解析。失败则结束，成功则将用户信息存入context向下一个handler传递
- router.go 中使用自定义的 jwt 中间件

### docker
项目跟目录下编写 Dockerfile 文件
```dockerfile
FROM golang:1.18

# 启用Go Modules依赖管理模式
RUN go env -w GO111MODULE=on

# 配置代理服务器地址， direct 则表示直接从原始的远程版本控制库拉取代码
RUN go env -w GOPROXY=https://goproxy.cn,direct

MAINTAINER tyc "1573496757@qq.com"

WORKDIR /home/app

# 将当前目录下的所有文件复制到容器内的/home/app目录下
COPY . .

# 将静态资源和日志文件挂载
VOLUME ["/home/app/storage"]

# go build命令来构建我们的Go应用程序
RUN go build -mod=mod main.go

# 暴露端口
EXPOSE 8001

ENTRYPOINT ["./main"]
```

打镜像： docker build -t go-web .
运行需配置个人mysql数据库信息例如
```markdown
DATABASE_HOST (默认值：127.0.0.1:3306)
DATABASE_USERNAME (默认值：root)
DATABASE_PASSWORD (默认值：123456)
DATABASE_DBNAME (默认值：t2 数据库名称)
```


	