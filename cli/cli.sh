#!/bin/bash

# https://github.com/go-micro/cli

go install github.com/go-micro/cli/cmd/go-micro@latest

# 创建微服务
go-micro new service github.com/luorufoeng/go-micro-example/cli/helloworld
cd helloworld
make init proto update tidy


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