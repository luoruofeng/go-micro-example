## 该client是用于测试myservice是否可调用。是成功启动consul和myservice的后续操作。

```shell
# 可以这样生成client，然后把myservice中的proto文件夹复制到这个client项目下使用。
go-micro new client --grpc myservice
```

- 如果没有通过docker来启动myservice则可以直接通过*go run .*来测试


- 下方依赖了myservice，是为引入myservice的proto。myservice已经依赖了go-micro，所以无须导入。
```shell
go mod init github.com/luoruofeng/go-micro-example/consul/client
go get github.com/luoruofeng/go-micro-example/consul/myservice
```

- 下方的consul_addr环境变量指的是consul的docker地址，如果是在同一个机器上做实验，请确保client和consul是在同一个network下。

- 在实际生产环境下 需要吧consul的8500端口给映射出来。因为client、consul，server都是在不同的机器上，他们互相交互是通过各自宿主机的ip地址。

```shell
#构建
docker build -t myclient:latest .

#运行（请先确保 myservice 和consul 已成功运行。下方的172.23.0.5是consul的容器ip。）
docker run --rm  --network=datacenter-deploy-secure_consul -e consul_addr=172.23.0.5:8500 myclient:latest
```
