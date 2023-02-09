如果想做负载均衡的服务端访问需要使用RoundRobin


```go 
selector := selector.NewSelector(
    selector.SetStrategy(selector.RoundRobin),
    selector.Registry(consulRegistry),
)

micro.Selector(selector)
```