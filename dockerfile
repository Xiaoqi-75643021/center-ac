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

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o center-ac

# 使用scratch作为最终镜像
FROM scratch

# 从builder镜像中复制构建的可执行文件到scratch镜像中
COPY --from=builder /app/center-ac .

# 复制config.json文件到scratch镜像中
COPY --from=builder /app/config.json .

# 运行应用程序
CMD ["./center-ac"]
