# devops 运维平台后端

主要实现运维平台后端 API 接口，包括服务管理、监控、配置中心等功能。

> 监控依赖于 prometheus

## 打包

```bash
# 构建
./build.sh

# 打包 docker 镜像
./docker-build.sh
```

## 框架

> 主要包括 api | pkg

### api

config | controllers | models | routers | api.go

#### config

> 服务配置初始化  
> 先加载默认配置，再读取配置文件，再读取环境变量，合并配置。

#### controllers

> controllers 主要包括业务逻辑处理。  
> 每一个业务模块对应一个 controller，使用注解路由。  
> base.go 实现路由拦截，包括一些控制器基础公用方法。
> auth.go 实现 jwt 认证路由拦截。

#### 数据模型 models

> models 主要包括基础业务数据处理  
> 每一个业务模块对应一个 model，并于 controllers 一一对应

#### 路由 routers

> routers/router.go 是路由文件，已于 gin 实现核心路由功能。

#### 主服务 api.go

> api.go 服务程序住文件，服务初始化全过程。

### pkg

auth | db | fileutil | resp | trans

#### auth

> 实现 jwt 认证加密算法。

#### db

> 实现数据库初始化的基础库。

#### fileutil

> 文件操作基础库。

#### resp

> 控制器返回 http json 数据的基础库。

#### trans

> gin 国际化基础库，base.go 里面会调用。
