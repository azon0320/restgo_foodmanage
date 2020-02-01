# Food Transport Rest Server
##### 食品运输管理(智能合约)网络，Fabric SDK 与 GIN 混合交互的后端服务器

这个服务器用在区块链 Hyperledger Fabric v1.0.0 的网络上，它通过Fabric SDK与区块链交互，用GIN与浏览器交互

## 兼容的测试环境与版本
* Golang (1.10.1)(没有go mod)
* Hyperledger Fabric SDK Go (79343ba)(用依赖管理工具 dep)
```
 git clone https://github.com/hyperledger/fabric-sdk-go
 cd fabric-sdk-go
 git checkout 79343ba
 # This sdk version uses dep
 dep ensure
```
* GIN (v1.3.0)(没有go mod)(用govendor解决依赖)
```
 git clone https://github.com/gin-gonic/gin
 cd gin
 git checkout v1.3.0
 govendor sync
```

## 快速部署
配置区块链网络需要配置文件，这个配置文件可以在Fabric SDK Go库`github.com/hyperledger/fabric-sdk-go/pkg/core/config/testdata/`里面可以找到

同时还需要配置通道名(Channel)、组织名(Organization)、证书(crypto-config)等

最快的配置指引在`/test/gin_integration_test.go`
``` shell script
 echo "编辑integration配置，SDK配置..."
 echo "运行快速部署"
 cd $GOPATH/src/github.com/dormao/restgo_foodmanage/test
 go test
 
 echo "关闭服务器"
 ^C
```

## 许可
[MIT](https://opensource.org/licenses/MIT)

## 其他语言
[English](./README.md)
