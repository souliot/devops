# DevOps developed by golang

## 部署方式

```shell
# 1.解压缩包 devops.tar.gz 得到以下文件
tar -zxvf devops.tar.gz
# .
# ├── devops-images.tar.gz
# ├── docker-compose.yml
# ├── load-images.sh
# ├── prom
# │   ├── conf
# │   │   └── prometheus.yml
# │   └── data
# └── save-images.sh

# 2.加载镜像文件
./load-images.sh

# 3.修改配置文件 docker-compose.yml
vim docker-compose.yml

# version: '3.7'

# services:
#   devops-ui:
#     container_name: devops-ui
#     hostname: devops-ui
#     image: devops-ui:5.1.2.0
#     ports:
#       - 38080:80     // 运维平台访问端口 38080（外部端口:容器内端口）--容器内端口固定
#     restart: always
#     depends_on:
#       - devops
#     ulimits:
#       nofile:
#         soft: 262144
#         hard: 262144
#     networks:
#       darwin:
#         aliases:
#           - devops-ui
#     logging:
#       driver: json-file
#       options:
#         max-file: '7'
#         max-size: 100m

#   devops:
#     container_name: devops
#     hostname: devops
#     image: devops:5.1.2.0
#     restart: always
#     ulimits:
#       nofile:
#         soft: 262144
#         hard: 262144
#     environment:
#       PROMADDRESS: prom:9090   // prometheus 地址（不用改）
#       DBHOST: 192.168.0.4:27017  // 本身存储数据的 mongo 地址；多个用 ; 隔开
#       DBNAME: Darwin-DevOps    //  数据库名称
#       DBUSER: ''    // 数据库用户名
#       DBPASSWORD: ''  // 数据库密码
#     volumes:
#       - ./prom/conf:/app/prom
#     networks:
#       darwin:
#         aliases:
#           - devops
#     logging:
#       driver: json-file
#       options:
#         max-file: '7'
#         max-size: 100m

#   prom:
#     container_name: prom
#     hostname: prom
#     image: prom/prometheus:v2.29.2
#     privileged: true
#     ports:
#       - 39090:9090   // prometheus 对外端口
#     volumes:
#       - ./prom/conf:/etc/prometheus
#       - ./prom/data:/prometheus
#       - /etc/localtime:/etc/localtime
#     command:
#       - '--web.enable-lifecycle'
#       - '--config.file=/etc/prometheus/prometheus.yml'
#       - '--storage.tsdb.path=/prometheus'
#       - '--web.console.libraries=/usr/share/prometheus/console_libraries'
#       - '--web.console.templates=/usr/share/prometheus/consoles'
#     restart: always
#     networks:
#       darwin:
#         aliases:
#           - prom
#     ulimits:
#       nofile:
#         soft: 262144
#         hard: 262144
#     logging:
#       driver: json-file
#       options:
#         max-file: '7'
#         max-size: 100m

# networks:
#   darwin:
#     external: true
```

## 问题

### 1、docker 网络不存在

```shell
# 这个是因为 这边没有直接创建网络，而是采用已有的网络"darwin",执行以下命令创建网络：

docker network create darwin
```
