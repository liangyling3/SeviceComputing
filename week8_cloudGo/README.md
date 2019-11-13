

## 项目地址
[GitHub项目地址](https://github.com/liangyling3/SeviceComputing/tree/master/week8_cloudGo)

##  概述
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）
## 任务基本要求
编程web应用程序cloudgo-io。请在项目 README.MD 给出完成任务的证据！ 基本要求：

- 支持静态文件服务
- 支持简单 js 访问
- 提交表单，并输出一个表格
- 对 /unknown 给出开发中的提示，返回码 5xx


## 测试结果
### 运行代码
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113195036784.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5neWxpbmcz,size_16,color_FFFFFF,t_70)
### 静态文件服务
assets文件夹下不存在index.html时：
- localhost:8080/static/

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113194659530.png)

添加index.html：
- localhost:8080/static/

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113194425746.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5neWxpbmcz,size_16,color_FFFFFF,t_70)
### 简单 js 访问
- localhost:8080/static/js/hello.js

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113194511620.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5neWxpbmcz,size_16,color_FFFFFF,t_70)
### 提交表单，并输出一个表格
- localhost:8080

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113194214918.png)
- localhost:8080/login

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113194310706.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5neWxpbmcz,size_16,color_FFFFFF,t_70)

### 对 /unknown 给出开发中的提示

-  localhost:8080/unknown

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113194821957.png)

- localhost:8080/sss

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113194845881.png)
## 注意事项
#### 相关包的安装
实验前需要使用 go get 命令安装相关包
```go
go get -u github.com/codegangsta/negroni
go get -u github.com/gorilla/mux
go get -u github.com/unrolled/render
```
#### 执行目录
必须在main.go目录下执行`go run`，而不能直接在其他目录下执行`go run '绝对路径'`，否则在访问时会出现`404 page not found`.

#### StripPrefix的使用
server.go中，需要使用语句
```go
mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
```
而不能直接使用
```go
mx.PathPrefix("/static").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
```
这是因为StripPrefix将访问时url中的"/static/“前缀去掉后再交给`http.FileServer(http.Dir(webRoot+"/assets/"))`处理，也就能访问到assets目录的路径，如果不使用StripPrefix而直接使用以下代码，则会出现404 page not found错误，因为该服务器目录中并没有”/static/"这个路径。
