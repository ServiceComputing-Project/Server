# Server——服务端项目

服务计算课程作业：https://github.com/ServiceComputing-Project

## 安装与运行

### 快速运行

```
git clone https://github.com/ServiceComputing-Project
```

若要部署，请修改关键信息[jwt、用户名、密码等]
数据库自带文章内容请在部署前删除

###  运行服务

- 配置```conf.toml```数据库信息
- 还原 ```data```目录下 ```db.sql``` 数据库
  数据库自带文章内容请在部署前删除
- 安装依赖
- 安装 swag   
  ```go get -u github.com/swaggo/swag/cmd/swag```
- 运行```swag init ```生成api文档
- 运行后台 ```go run```  

###  运行后台

- 安装依赖 ``` npm install ```
- 开发运行 ``` npm run serve ```
- 浏览器打开 [http://127.0.0.1:8080/](http://127.0.0.1:8080/)
- 发布 ```npm run build ``` 会自动发布到 ```dist```目录下面
- 友链里面第一个为后台登陆地址默认用户名```zxysilent```,密码```zxyslt```，可自行数据库修改

### 评论配置

- 配置项目 opts(表).comment(值) 
- 配置说明 https://github.com/gitalk/gitalk

