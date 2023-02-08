## 配置可以从consul预先设置了的kv中加载，也可以重文件，从内存(变量)加载。

目前go-micro只支持consul加载json格式，如果要支持yaml，需要添加新的yaml的encoder,参考：https://github.com/go-micro/go-micro/issues/614

```shell
go get github.com/go-micro/plugins/v4/config/encoder/yaml
```