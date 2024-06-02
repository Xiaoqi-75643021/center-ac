# 使用官方Go镜像作为构建环境
FROM golang:1.22.3 AS builder

# 设置国内的Golang代理
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 将go.mod和go.sum复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将项目的所有文件复制到工作目录
COPY . .

# 构建Go应用程序
RUN go build -o main .

# 暴露应用程序端口
EXPOSE 8080

# 运行Go应用程序
CMD ["./main"]
