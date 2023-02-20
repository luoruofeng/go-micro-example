## 如何使用go-micro/cli
refer to:https://github.com/go-micro/cli

```shell
go install github.com/go-micro/cli/cmd/go-micro@latest

# 创建微服务
go-micro new service github.com/luorufoeng/go-micro-example/cli/helloworld
cd helloworld
make init proto update tidy

# 也可以加上相关的参数，如：
# go-micro new service --grpc --health --jaeger github.com/luorufoeng/go-micro-example/cli/helloworld
# cd helloworld


# 指定端口
# srv.Init方法中添加micro.Address(":8080"),

# 运行微服务
go-micro run
```
----
## 如何调用服务
```shell
# client调用server(linux下调用)
go-micro call helloworld Helloworld.Call '{"name": "Luo"}'
```

----
## 如何生成Client
如果是带参数：--grpc生成的server端，不能直接通过go-micro call来调用，必须要生成client来调用，如下所示：
```shell
# 命令行生成client:
go-micro new client --grpc github.com/luorufoeng/go-micro-example/cli/helloworld
# 要把server的proto文件夹复制到项目中：
cp -r ../helloworld/proto ./
#如果go mod use .失败，需要修改go.mod的replace,但是修改了这个后导包都需要改名
make update tidy
go run .

```
----
## 构建容器
Dockerfile中添加下面代码
```shell
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add --no-cache git
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
```
构建docker镜像，并启动
```shell
make docker
docker run -p 8080:8080 helloworld:latest
```
## go-micro中的proto生成的xxx.pb.micro.go文件
----
1. 这个文件中的NewxxxService方法返回的是client调用service时候使用的stub。
2. stub中的方法是client调用service时使用xxxService接口中定义的方法。
3. 这个文件中的xxxHandler是当前service的handler中需要实现的接口。
