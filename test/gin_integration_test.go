package test

import (
	"github.com/dormao/restgo_foodmanage/router"
	"github.com/dormao/restgo_foodmanage/sdk"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"log"
	"net/http"
	"testing"
)

/*
 * This test runs a server with full feature routes
 *
 * You must do some work first
 * 1. Launch E2E network
 * 2. ConfigFile: $PROJECT_ROOT/sdk/test/configs/config_test_e2e.yaml
 * 3. Specify channel name, username, org name, chaincode name
 *
 * The fabsdk config examples are in github.com/hyperledger/fabric-sdk-go/pkg/core/config/testdata
 * Copy and specify it to sdk
 */
func TestIntegration(t *testing.T) {
	const (
		IntegrationConfigFile     = "../sdk/test/configs/config_test_e2e.yaml"
		IntegrationChannelName    = "mychannel"
		IntegrationUserName       = "User1"
		IntegrationOrgName        = "org1"
		IntegrationChaincodeName  = "mycc"
		IntegrationStaticFilePath = "../resource"
	)
	log.Println("Init Default Fabric SDK Client")
	err := sdk.GlobalInitClient(
		config.FromFile(IntegrationConfigFile), IntegrationChannelName,
		IntegrationUserName, IntegrationOrgName, IntegrationChaincodeName)
	if err != nil {
		log.Println("Can not start client")
		log.Fatal(err)
		return
	}
	log.Println("Execute ping to network and wait...")
	err = sdk.ExecutePing()
	if err != nil {
		log.Fatal("Initial execute failed" + err.Error())
		return
	}
	log.Println("Start GIN Server")
	engine := gin.Default()
	engine.StaticFS("/static", http.Dir(IntegrationStaticFilePath))
	router.RegisterRouter(engine)
	engine.Run("0.0.0.0:8080")
}
