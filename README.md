<h3 align="center">About NigTusg.</h3>
<!-- <h3 align="center">About NigTusg.</h3> 是HTML语法，用于创建居中的三级标题。-->
<div style="display:flex" height="auto" width="auto"> <!-- 第一个个div表示为第一个大的区间块，用于存放代码-->
    <!-- <div style="display:flex" height="auto" width="auto"> 是一个HTML标签，用于创建一个HTML元素，这个元素被称为 “div”。“div” 是 “division” 的缩写，意味着它可以用来划分页面的不同部分。
         在这个标签中，style="display:flex" 是一个CSS属性，它设置了这个div的布局为弹性盒子（flexbox）。弹性盒子是一种用于在页面上对元素进行布局、对齐和分配空间的工具。
         height="auto" 和 width="auto" 是HTML属性，它们设置了这个div的高度和宽度自动适应其内容的大小。
    -->
    <div>
        <div style="flex:1">
            <img src="https://media.giphy.com/media/WUlplcMpOCEmTGBtBW/giphy.gif" width="40"></div>
            Hello！茫茫人海，感谢相遇 🔆  
            <img align="right" alt="GIF" src="https://user-images.githubusercontent.com/105040964/205486218-b8a4f47d-6a8e-420e-bba6-0ddce8c4deac.gif" width="430" height="100%" /> 
         </div>
    <!-- 这是一个HTML标签，用于在网页上插入一张图片 其中，src这部分定义了图片的来源地址，即为图片的URL
         alt="GIF" 这部分为图片的替代文本，如果图片显示异常，就会显示这个文本
         align="right"：这部分设置了图片的对齐方式，使图片在其容器中右对齐。
         width="430"：这部分设置了图片的宽度为430像素。
         height="100%"：这部分设置了图片的高度为其容器的100%。
    -->
    <!-- 在这第一个嵌套的子区间代码块中，分别嵌入了两个div，一个是简介，一个是动图-->
       
</br>
        <div style="flex:1" >
            <ul >
              <li>🎇 喜欢分享编程知识，更欢迎来交流学习。 </li>
              <li>⛅ 希望有一天能够成为很厉害的架构师。 </li>
              <li>🌀 我的邮箱 : XuZh01050@gmail.com </li>
              <li>🌊 一名计算机科学与技术的本科生 </li>
              <li>🌴 想要去杭州工作  </li>
            </ul>  <!-- <ul> 和 <li> 是用来创建无序列表的HTML标签 -->
        </div>      
</div> 
    <br>
    </br>



## 上手指南

### 启动服务

将`config.yaml`中所有host改为本机地址后输入

```bash
docker-compose up
```

即可通过docker快速启动部署服务及相关依赖服务

### 相关环境

- **golang**>= 1.18
- **mysql**>=8.0：数据库
- **redis**>=7.0.0：缓存
- **minio**：对象存储
- **ffmpeg**：获取视频封面

## 技术选型

<img src="https://img-blog.csdnimg.cn/0d5cabef362d4f71a5051b44596745c1.png" width="50%" height="50%" >

## 实现功能

|    功能    |                             说明                             |
| :--------: | :----------------------------------------------------------: |
|  基础功能  |      视频feed流、视频投稿，个人信息、用户登录、用户注册      |
| 扩展功能一 | 视频点赞/取消点赞，点赞列表；用户评论/删除评论，视频评论列表 |
| 扩展功能二 |            用户关注/取关；用户关注列表、粉丝列表             |


## 目录结构 

```C:.
├─.idea
│  └─dataSources
│      └─b1278ade-bf74-413c-a427-1bd4fc0f0dcc
│          └─storage_v2
│              └─_src_
│                  └─schema
├─common
├─config
├─controller
├─log
├─minioStore
├─proto
│  ├─pkg
│  └─proto
├─repository
├─response
├─routes
├─service
└─util
PS C:\Users\0\Desktop\TikTokLite-main> tree /f
卷 Windows 的文件夹 PATH 列表
卷序列号为 04BD-DE10
C:.
│  config.yaml
│  docker-compose.yml
│  Dockerfile
│  go.mod
│  go.sum
│  main.go
│  README.md
│  redis.conf
│  TikTokLite.sql
│  wait-for.sh
│
├─.idea
│  │  .gitignore
│  │  dataSources.local.xml
│  │  dataSources.xml
│  │  modules.xml
│  │  TikTokLite.iml
│  │  vcs.xml
│  │  workspace.xml
│  │
│  └─dataSources
│      │  b1278ade-bf74-413c-a427-1bd4fc0f0dcc.xml
│      │
│      └─b1278ade-bf74-413c-a427-1bd4fc0f0dcc
│          └─storage_v2
│              └─_src_
│                  └─schema
│                          information_schema.FNRwLQ.meta
│                          information_schema.FNRwLQ.zip
│                          mysql.osA4Bg.meta
│                          performance_schema.kIw0nw.meta
│                          sys.zb4BAA.meta
│                          TikTokLite.yLxFgg.meta
│                          TikTokLite.yLxFgg.zip
│
├─common
│      AuthMiddleware.go
│      cache.go
│      dbInit.go
│
├─config
│      config.go
│
├─controller
│      commentController.go
│      favortiteController.go
│      feedController.go
│      publishController.go
│      relationController.go
│      userController.go
│
├─log
│      log.go
│
├─minioStore
│      minioClient.go
│
├─proto
│  ├─pkg
│  │      comment.pb.go
│  │      favorite.pb.go
│  │      feed.pb.go
│  │      login.pb.go
│  │      publish.pb.go
│  │      register.pb.go
│  │      relation.pb.go
│  │      user.pb.go
│  │
│  └─proto
│          comment.proto
│          favorite.proto
│          feed.proto
│          login.proto
│          publish.proto
│          register.proto
│          relation.proto
│          user.proto
│
├─repository
│      commentModel.go
│      favoriteModel.go
│      relationModel.go
│      userModel.go
│      videoModel.go
│
├─response
│      response.go
│
├─routes
│      comment.go
│      favorite.go
│      publish.go
│      relation.go
│      routes.go
│      user.go
│
├─service
│      commentService.go
│      favoriteService.go
│      feedService.go
│      publishService.go
│      relationService.go
│      userService.go
│
└─util
        util.go
```

- `common`：中间件、数据库初始化
- `config`： 读取配置
- `controller`：视图层，处理前端消息
- `log`：zap日志组件进行封装
- `minioStore`：对象存储服务，生成视频对外访问连接
- `proto`：前端消息结构体，由`protobuf`文件自动生成
- `repository`：数据层，直接对数据库进行操作
- `response`：对返回消息进行封装
- `routes`：路由层
- `service`：逻辑层，执行业务操作，从数据层获取数据，封装后返回试图层
- `uitl`：工具函数
- `TikTokLite.sql`：数据库建表文件 
- `config.yaml`：配置文件
- `redis.conf`：redis配置文件
- `main.go`：服务入口

## 开发整体设计

### 整体架构图

<img src="https://img-blog.csdnimg.cn/cc6070c6e54a40dc95ea9b34cd855aa8.png"  width="65%" height="65%"  >



### 数据库设计

<img src="https://img-blog.csdnimg.cn/be4524a1a81e4a31a6699ea03c3466f2.png" width="65%" height="65%"  >

## 优化

### 1. 安全

1. 引入JWT，进行`全局Token管理`，高效管理用户Token，并且设置过期时间。
2. Redis引入`redsync锁`机制，防止俩个线程同时修改用户信息(例如关注)
3. Redis引入`事务`机制，防止多表操作时，只修改一张表。最终导致失败。
4. 使用参数占位符来构造SQL语句，不使用字符串拼接，`避免SQL注入`
5. 用户密码进行`MD5加密`处理，返回用户基本信息时进行`脱敏`。
6. 实现`鉴权中间件`，将鉴权和实际业务分离，对不同的接口设置不同的访问权限
7. 使用`docker`整合所有相关依赖服务，便于用户快速部署服务，使用wait-for确保其他依赖服务启动后再启动后端服务

### 2. 性能
1. 根据实际业务,Querry语句的需求，合理`设置相关索引`，保证索引高命中
2. 引入`Redis`作为中间件，用来实现对象缓存，提升响应速度，减少IO操作，减少服务器压力
3. 通过`Minio`自己搭建对象存储，来存储上传视频，并且将上传的视频生成URL，并将URL放在数据库中，避免存储冗余。
4. 通过`pprof`进行性能测试，引入缓存与无缓存之间的性能

### 3. 项目维护

1. 项目`Git协同`,严格遵循成员PR->Review->Merge三步走流程，避免错误代码扩散到其他成员库
2. `多次迭代目录结构`。目录结构清晰，配置单元、日志单元、各模块单元条理分明

3. 文档管理。修改代码前后，要随时记录文档，跟进开发流程，字最后测试出现问题时，可以找到具体负责人快速解决


## 未来展望

- 分布式

  利用grpc作为分布式框架，etcd或zookeeper作为注册中心，将五个模块分别布置到不同的服务器上，通过RPC远程调用的方式，来调用相关的模块的方法，做到分布式处理与解耦

<img src="https://img-blog.csdnimg.cn/f8e7445378f04f8ba77772a774c2afc0.png"  width="65%" height="65%" >



## 版本控制
- 该项目使用Git进行版本管理。您可以在repository参看当前可用版本。


