
```shell
#启动consul
cd datacenter-deply-secure
docker-compose up -d
```
- 访问172.0.0.1:8500 查看consul面板

```
#创建微服务
go-micro new service github.com/luorufoeng/go-micro-example/consul/myservice
```

- *proto/myservice.proto*已经做了修改删除了不必要的方法

```shell
#创建proto的对应文件和下载依赖
cd myservice
make init proto update tidy

#依赖导入
go get github.com/hashicorp/consul/api
go get github.com/go-micro/plugins/v4/registry/consul

#启动 
go run .
```
- 如果是直接启动服务，则可以运行*client文件夹*中的go run .访问,来测试微服务是否成功。


## 构建容器镜像

- *Dockerfile*中添加下面代码
```
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add --no-cache git
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
```

```shell
#构建
make docker
#运行(使用的是172.23.0.5是sconsul1的容器ip。consul和myservice都在同一个network中，所以可以使用。consul目前是通过docker运行的，端口映射到宿主机。也可以使用宿主机的IP。)
docker run  --network=datacenter-deploy-secure_consul -e myservice_port=8080 -e consul_addr=172.23.0.5:8500 -p 8080:8080 myservice:latest
```

- 运行后请确保service和consul都是使用docker，并且在同一个网络下。可以互相ping通。

- 如果是容器启动服务，需要把*client文件夹*中的main.go也做成容器镜像运行，并且保证，server和client的容器可以互相ping通（使用docker的默认桥连接网络或者使用同一个network）