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

# 构建静态链接的Go应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用scratch作为运行环境
FROM scratch

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制构建的二进制文件
COPY --from=builder /app/main .

# 运行Go应用程序
CMD ["./main"]
