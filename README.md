# go-blog
## 项目配置
mv conf/config.example.toml conf/config.toml
配置config.toml

## 安装npm
yum install npm

## 前端资源安装及打包
cd /project
npm install
npm run build

## 启动服务
go run main.go
或者
go build -o go-blog
./go-blog