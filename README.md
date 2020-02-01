# Food Transport Rest Server
##### FoodTransport Manager RestAPI Server for Hyperledger Fabric v1.0.0 network

This is the http server for hyperledger network, it works with Fabric SDK Go, invoke and query chaincode functions by Fabric SDK Client.

## Compatiable Libraries and Environments
* Golang (1.10.1)(no go mod)
* Hyperledger Fabric SDK Go (79343ba)(use dep)
```
 git clone https://github.com/hyperledger/fabric-sdk-go
 cd fabric-sdk-go
 git checkout 79343ba
 # This sdk version uses dep
 dep ensure
```
* GIN (v1.3.0)(no go mod)(use govendor)
```
 git clone https://github.com/gin-gonic/gin
 cd gin
 git checkout v1.3.0
 govendor sync
```

## Quick Launch
Launch the server must config some files first, the Fabric SDK config files can be found in directory `github.com/hyperledger/fabric-sdk-go/pkg/core/config/testdata/`

It must be configured with Channel Name, Organization Name, crypto-path, etc..

The fastest launch code is in `/test/gin_integration_test.go`

``` shell script
 echo "Edit the integration code with network config, crypto config..."
 echo "Launch Quick Launch"
 cd $GOPATH/src/github.com/dormao/restgo_foodmanage/test
 go test
 
 echo "Stop the server"
 ^C
```

## License
[MIT](https://opensource.org/licenses/MIT)

## Other Language
[Chinese](./README_cn.md)
