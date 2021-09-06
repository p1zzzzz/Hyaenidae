# Hyaenidae

go和vue基础学习之作

markdown协作平台

## 需求

* markdown协作
* 具备拓展性
* 角色权限管理
* 前后端分离

## 部署

### 后端
```bash
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server main.go 

# 运行二进制
./server 
```

### 前端
```bash
# 进入web文件夹
cd web

# 安装依赖
npm install

# 启动web项目
npm run serve
```

默认用户名/密码 admin/admin123


## 技术选型
* 前端：用基于 [Vue](https://vuejs.org/) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。
*  后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
*  数据库：采用`MySql`，使用 [gorm](http://gorm.cn/) 实现对数据库的基本操作。
*  配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
*  日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

## 目录结构
```bash
├── service
│   ├── api #api接口
│   │   └── v1 
│   ├── config #配置相关
│   ├── core #核心
│   ├── global #全局变量
│   ├── initialize #初始化
│   ├── middleware #中间件
│   ├── model #模型层
│   │   ├── request 
│   │   └── response
│   ├── resource #静态资源
│   ├── router #路由
│   ├── service #服务
│   ├── source #source
│   └── utils #工具包
│       └── upload 
└── web
    ├── public
    └── src
        ├── api #向后台发送ajax的封装层
        ├── assets #静态文件
        ├── components #组件
        ├── directive #前端路由
        ├── mixins #混入
        ├── router #前端路由
        ├── store #vuex状态管理
        ├── style #样式
        ├── utils #前端工具包
        └── view #前端页面
```
## todo

- [ ] 集成更多在线工具，如默认口令查询

## 参考
gin-vue-admin: https://github.com/flipped-aurora/gin-vue-admin
elementUI: https://element.eleme.io/#/zh-CN
