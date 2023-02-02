#!/bin/bash

# https://github.com/go-micro/cli

go install github.com/go-micro/cli/cmd/go-micro@latest

# 创建微服务
go-micro new service github.com/luorufoeng/go-micro-example/cli/helloworld
cd helloworld
make init proto update tidy

# 也可以加上相关的参数，如：
# go-micro new service --grpc --health github.com/luorufoeng/go-micro-example/cli/helloworld
# cd helloworld

# client也可以通过命令行生成，要把server的proto文件夹复制到项目中：
# go-micro new client --grpc github.com/luorufoeng/go-micro-example/cli/helloworld

# 指定端口
# srv.Init方法中添加micro.Address(":8080"),

# 运行微服务
go-micro run

# 调用grpc
go-micro call helloworld Helloworld.Call '{"name": "John"}'

# 构建容器
# 先启动docker
# Dockerfile中添加下面代码
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
# RUN apk update && apk add --no-cache git
# RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct

make docker
docker run -p 8080:8080 helloworld:latest