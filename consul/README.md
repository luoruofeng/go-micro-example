Launch consul docker befor running program

reference https://github.com/go-micro/plugins/blob/main/v4/registry/consul/registry_test.go

```shell
cd datacenter-deply-secure
docker-compose up -d

```

```go
cd code 
go get github.com/hashicorp/consul/api
go get github.com/go-micro/plugins/v4/registry/consul
```


```shell
cd code 
go run .
```