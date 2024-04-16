# 第一阶段：使用golang镜像进行Go构建
FROM golang:1.21.6 AS build-go

# 设置工作目录
WORKDIR /app

# 拷贝go.mod和go.sum文件以下载依赖
COPY go.mod go.sum ./

# 下载Go依赖
RUN go mod download

# 拷贝项目文件到工作目录
COPY . .

# 构建Go应用
RUN CGO_ENABLED=0 GOOS=linux go build -o go-blog ./main.go

# 第二阶段：使用node镜像进行npm构建
FROM node:21.7.1 AS build-npm

# 设置工作目录
WORKDIR /app

# 拷贝package.json和package-lock.json文件
COPY package*.json webpack.config.js ./

# 安装npm依赖
RUN npm ci --only=production

# 拷贝前端代码到工作目录
COPY ./static ./static

# 安装webpack-cli
RUN npm install -D webpack-cli

# 设置环境变量，添加./node_modules/.bin到PATH
ENV PATH /app/node_modules/.bin:$PATH

# 构建前端代码
RUN npm run build

# 第三阶段：使用alpine镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root

# 创建日志目录和文件
RUN mkdir -p logs && touch logs/logfile.log

# 拷贝配置文件和模板文件到容器
COPY conf/config.toml /root/conf/config.toml
COPY views /root/views

# 从第一阶段拷贝Go应用
COPY --from=build-go /app/go-blog .

# 从第二阶段的build-npm阶段拷贝Go应用和npm构建的静态文件
COPY --from=build-npm /app .

# 构建前端代码并查看输出目录
RUN ls -l ./static/dist/

# 设置运行命令
CMD ["./go-blog"]