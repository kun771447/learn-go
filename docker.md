

## 安装 Docker

### mac

https://www.docker.com/

## docker 命令

```shell
docker ps -a |grep elasticsearch
```


## 安装 Mysql


### 1. 拉取MySQL镜像:
在终端中运行以下命令拉取MySQL官方镜像：

```
docker pull mysql
```

### 2. 创建MySQL容器:
运行以下命令创建并启动一个MySQL容器：

```
docker run --name my-mysql-container -e MYSQL_ROOT_PASSWORD=your_password -d -p 3306:3306 mysql
```

--name my-mysql-container: 为容器指定一个名称，你可以自己取一个。
-e MYSQL_ROOT_PASSWORD=your_password: 设置MySQL root用户的密码，请将 your_password 替换为你想要设置的密码。
-d: 在后台运行容器。
-p 3306:3306: 将容器的3306端口映射到主机的3306端口。

### 3.验证MySQL容器是否运行:
运行以下命令查看运行中的容器：

```
docker ps
```

你应该能够看到 my-mysql-container 或你指定的容器名称。

### 4.使用MySQL客户端连接到MySQL容器:
你可以使用MySQL客户端连接到运行中的MySQL容器。例如，可以使用以下命令连接到MySQL容器：

```
docker exec -it my-mysql-container mysql -uroot -p
```