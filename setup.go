package main

import (
	"fmt"
	"github.com/dormao/restgo_foodmanage/router"
	"github.com/dormao/restgo_foodmanage/sdk"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"net/http"
)

var (
	ConfigFilePath    = ""
	SdkChannelName    = ""
	SdkChaincodeName  = ""
	SdkOrgName        = ""
	SdkUserName       = ""
	ResourceDirectory = "./resource"
	ServerAddr        = "0.0.0.0:80"
)

func main() {
	err := sdk.GlobalInitClient(
		config.FromFile(ConfigFilePath),
		SdkChannelName, SdkUserName, SdkOrgName, SdkChaincodeName)
	if err != nil {
		fmt.Println("error init fabsdk client : " + err.Error())
		return
	}
	err = sdk.ExecutePing()
	if err != nil {
		fmt.Println("fabsdk invoke fail : " + err.Error())
		return
	}
	engine := gin.Default()
	engine.StaticFS("/static", http.Dir(ResourceDirectory))
	router.RegisterRouter(engine)
	_ = engine.Run(ServerAddr)
}
