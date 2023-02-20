# 启动jaeger
```shell
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.42
```

# 创建项目
```shell
   github.com/micro/go-micro


   go-micro new service --grpc --health --jaeger aservice
   go-micro new service --grpc --health --jaeger bservice
   go-micro new service --grpc --health --jaeger cservice

   go work use ./jaeger/aservice
   go work use ./jaeger/bservice
   go work use ./jaeger/cservice

   go get github.com/go-micro/cli/debug/trace/jaeger
```


调用链路为(三个service的proto定义需要包含他们依赖的service的全部定义):
CService -->> BService -->> AService

http://127.0.0.1:16686/查询
----

## 创建client 测试
```shell
go-micro new client --grpc cservice
cd cservice-client
go work use .
cp -r ../cservice/proto ./
make update tidy
#修改 go.mod 中的replace和main.go中的NewAserviceService
go run .
```