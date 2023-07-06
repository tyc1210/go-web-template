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