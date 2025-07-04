# FuseOps-Backend

## 项目简介

fuseops-backend 是一个基于 Gin 框架的运维管理系统后端，支持用户管理、权限认证、JWT 登录、Swagger 文档等功能，适合中小型企业或团队的运维自动化场景。

## 目录结构

```
fuseops-backend/
├── internal/
│   ├── api/              # 路由注册
│   ├── handler/          # 业务接口实现
│   ├── model/            # 数据库模型
│   ├── repository/       # 数据库操作
│   ├── middleware/       # Gin 中间件
│   └── pkg/              # 工具包
├── config.yaml           # 配置文件
├── go.mod                # Go module 文件
├── go.sum                # Go 依赖校验文件
├── main.go               # 主程序入口
└── README.md             # 项目说明
```

## 快速开始

### 1. 克隆项目

```shell
git clone https://github.com/sreok/fuseops-backend.git
cd fuseops-backend
```

### 2. 配置数据库和JWT

编辑 `config.yaml`，填写数据库和JWT相关配置。

### 3. 安装依赖

```shell
go mod tidy
```

### 4. 生成Swagger文档

```shell
swag init -g main.go
```

### 5. 启动服务

```shell
go run main.go
```

### 6. 访问接口文档

浏览器访问 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## 常用命令

- 生成Swagger文档：`swag init -g main.go`
- 启动服务：`go run main.go`

## 相关技术

- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Viper](https://github.com/spf13/viper)
- [Swaggo](https://github.com/swaggo/swag)
- [JWT](https://github.com/golang-jwt/jwt)

## 提交记录

- 2025.7.4 - 更新

1. 添加用户多角色关联表，登录json返回用户多角色
2. 重构api与handler

- 2025.7.2 - 首次提交

1. 用户增删改查（CRUD）
2. JWT 登录认证
3. Bcrypt 密码加密
4. 角色与权限管理
5. 支持自定义表前缀
6. 配置集中管理（config.yaml）
7. Swagger 自动接口文档
