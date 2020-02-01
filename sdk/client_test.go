package sdk

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"testing"
)

var (
	testCfgFile      = "./test/configs/config_test_e2e.yaml"
	testChannelName  = "mychannel"
	testOrganization = "org1"
	testUserName     = "User1"

	testChaincodeID = "mycc"
	testFunction    = "reg_seller"
	testArgs        = [][]byte{
		[]byte("testPassword"),
	}
)

func createClient() (*fabsdk.FabricSDK, *channel.Client, error) {
	cfg := config.FromFile(testCfgFile)
	return SetupClient(cfg, testChannelName, testUserName, testOrganization)
}

func TestClientInit(t *testing.T) {
	sdk, _, err := createClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer sdk.Close()
	t.Log("Init Success")
}

func TestClientInvoke(t *testing.T) {
	fmt.Println("Begin invoke {Function: reg_seller Args: [testPassword]}")
	fmt.Println("Create Client")
	sdk, client, err := createClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer sdk.Close()
	fmt.Println("Execute chaincode")
	resp, err := client.Execute(channel.Request{
		ChaincodeID: testChaincodeID,
		Fcn:         testFunction,
		Args:        testArgs,
		// The transient map has secure data, which will not broadcast to some transaction
		// It will be also sent to chaincode to process
		// Chaincode use: GetTransientMap
		TransientMap: nil,
	})
	if err != nil {
		for _, r := range resp.Responses {
			t.Log(r.String())
		}
		t.Error(err)
		return
	}
	out := "Response Payload : " + string(resp.Payload)
	fmt.Println(out)
	t.Log(out)
}
