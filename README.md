<div align="center">

# server-status-server

_✨ 服务器状态 - 服务端/后端 ✨_

</div>

## 📖 介绍

服务器状态的服务端

## 💿 安装（还在推进）

### 从二进制文件安装

- 从 Release 下载对应平台的二进制文件

### Docker 安装

- 从 Docker Hub 下载镜像
  docker pull liteyuki/server-status-server

### 自行编译

- 你无法自行编译，因为应用包含私有依赖

## 🎉 使用

### 配置
默认是通过传入环境变量来配置的，你可以通过创建 `.env` 文件来或者直接传入环境变量来配置
```dotenv
PORT=8080
TOKEN=114514
```

## 📝 其他

### 开机启动

- 手动安装请自行配置service通过systemd启动

### 主机监控

- 请在需要监控的服务器主机上安装 [server-status-client](../server-status-client)