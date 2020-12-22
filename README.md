

# 客户端项目：Go API Server for swagger

## 参与人员

- [EmilyYoung9](https://github.com/orgs/ServiceComputing-Project/people/EmilyYoung9) 杨玲（18342115）
- [Annecle](https://github.com/orgs/ServiceComputing-Project/people/Annecle)  周圆（18342143）
- [beikenken](https://github.com/orgs/ServiceComputing-Project/people/beikenken) 张又方（）
- [utaZ ](https://github.com/orgs/ServiceComputing-Project/people/utaZ) 邹文睿（18342146）

## 程序包功能

实现简单的blog，含有注册登录功能、查看文章、查看评论及进行评论的功能。

## 软件架构

**go文件夹**：具体实现，包含comment、user、article三个部分的实现。

**test文件夹**：所有函数实现的对应测试文件。

**data文件夹**：从网络抓取的公开技术文章内容和相应评论信息。

后端分工：

EmilyYoung9 杨玲（18342115）

Annecle 周圆（18342143）

beikenken 张又方（）

utaZ 邹文睿（18342146）


## 安装与运行

执行如下命令进行安装：

```
# 安装 Bolt DB
go get -u github.com/boltdb/bolt

# 安装 JWT 模块
go get -u github.com/dgrijalva/jwt-go

# 安装执行http请求的路由和分发的第三方扩展包mux
go get -u github.com/gorilla/mux

# 安装项目
go get -v github.com/ServiceComputing-Project/Server
```

执行如下命令运行：

```
go run main.go
```

## 测试说明

* ### signin，记得保存 token

```shell
http://localhost:8080/simpleblog/user/signin?username=user5&password=pass5
```

* ### getArticleById

```
http://localhost:8080/simpleblog/user/article/1
```

* ### getArticles

```
http://localhost:8080/simpleblog/user/articles?page=1
```

* ### deleteArticle

```
http://localhost:8080/simpleblog/user/deleteArticle/1
```

* ### createComment 
  
  * 将之前登陆的 user 的 token 放在 Authorization 后，author 对应登陆的 user

```
curl -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU2OTE4MzcsImlhdCI6MTU3NTY4ODIzN30.4Uw5rYZPCxB7uXNrVtn69Tmsy-831CgPnF8e555z-ko" http://localhost:8080/simpleblog/user/article/2/comment -X POST -d '{"content":"new content3","author":"user5"}'
```

* ### getComments

```
http://localhost:8080/simpleblog/user/article/2/comments
```



- "token":"eyJhbGciOiJIUzI1NiIsInR5c
                      CI6IkpXVCJ9.eyJleHAiOjE1NzU2OTQxOTY
                      sImlhdCI6MTU3NTY5MDU5Nn0.ZOzLig7pRA
                      tKTKlhR4e_uJlCEc5Ehn5FYlGMrQIouJQ"
- curl -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU2OTQxOTYsImlhdCI6MTU3NTY5MDU5Nn0.ZOzLig7pRAtKTKlhR4e_uJlCEc5Ehn5FYlGMrQIouJQ" http://localhost:8080/simpleblog/user/article/2/comment -X POST -d '{"content":"new content3","author":"user5"}'
- curl -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU2OTQxOTYsImlhdCI6MTU3NTY5MDU5Nn0.ZOzLig7pRAtKTKlhR4e_uJlCEc5Ehn5FYlGMrQIouJQ" http://localhost:8080/simpleblog/user/article/2/comment -X POST -d '{"content":"new content3","author":"user5"}'
