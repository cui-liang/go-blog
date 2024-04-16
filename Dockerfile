# 第一阶段：使用golang镜像进行Go构建
FROM golang:1.21.6 AS build-go

# 设置工作目录
WORKDIR /app

# 拷贝所有源代码和依赖文件
COPY . .

# 下载Go依赖
RUN go mod download

# 构建Go应用
RUN CGO_ENABLED=0 GOOS=linux go build -o go-blog ./main.go

# 第二阶段：使用node镜像进行npm构建
FROM node:21.7.1 AS build-npm

# 设置工作目录
WORKDIR /app

# 拷贝前端代码、依赖文件和Webpack配置文件
COPY ./static ./static
COPY package*.json webpack.config.js ./

# 安装npm依赖并构建前端代码
RUN npm ci --only=production --omit=dev && \
    npm install -D webpack-cli && \
    npm run build

# 第三阶段：使用alpine镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root

# 创建日志目录和文件
RUN mkdir -p logs && touch logs/logfile.log

# 拷贝配置文件和模板文件到容器
COPY conf/config.toml conf/config.toml
COPY views views

# 从第一阶段拷贝Go应用
COPY --from=build-go /app/go-blog .

# 从第二阶段的build-npm阶段拷贝Go应用和npm构建的静态文件
COPY --from=build-npm /app/static ./static

# 设置运行命令
CMD ["./go-blog"]