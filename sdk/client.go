package sdk

import (
	"errors"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var (
	defaultSdk           *fabsdk.FabricSDK = nil
	defaultClient        *channel.Client   = nil
	defaultChannelName                     = ""
	defaultChaincodeName                   = ""
)

func GetClient() *channel.Client {
	return defaultClient
}

func ClientExecute(fcn string, args [][]byte) (channel.Response, error) {
	if GetClient() == nil {
		return channel.Response{}, errors.New("No sdk client launched")
	}
	return GetClient().Execute(channel.Request{
		ChaincodeID:  defaultChaincodeName,
		Fcn:          fcn,
		Args:         args,
		TransientMap: nil,
	})
}

func ClientQuery(fcn string, args [][]byte) (channel.Response, error) {
	if GetClient() != nil {
		return channel.Response{}, errors.New("No sdk client launched")
	}
	return GetClient().Query(channel.Request{
		ChaincodeID:  defaultChaincodeName,
		Fcn:          fcn,
		Args:         args,
		TransientMap: nil,
	})
}

func ExecutePing() error {
	resp, err := ClientExecute(models.UnAuthPing, [][]byte{})
	if err != nil {
		return err
	}
	println(string(resp.Payload))
	return nil
}

func QueryPing() (channel.Response, error) {
	return GetClient().Query(channel.Request{
		ChaincodeID:  defaultChaincodeName,
		Fcn:          models.UnAuthPing,
		Args:         [][]byte{},
		TransientMap: nil,
	})
}

func GlobalInitClient(cfg core.ConfigProvider, chname string, user string, org string, ccname string) error {
	if defaultSdk != nil {
		return errors.New("fabric sdk is running")
	}
	sdk, client, err := SetupClient(cfg, chname, user, org)
	if err != nil {
		return err
	}
	defaultClient = client
	defaultSdk = sdk
	defaultChannelName = chname
	defaultChaincodeName = ccname
	return nil
}

func GlobalCloseClient() {
	if defaultClient != nil {
		defaultClient = nil
	}
	if defaultSdk != nil {
		defaultSdk.Close()
		defaultSdk = nil
	}
}

func SetupClient(cfg core.ConfigProvider, chname string, user string, org string) (*fabsdk.FabricSDK, *channel.Client, error) {
	sdk, err := fabsdk.New(cfg)
	if err != nil {
		return nil, nil, err
	}
	chProvider := sdk.ChannelContext(chname,
		fabsdk.WithOrg(org), fabsdk.WithUser(user))
	client, err := channel.New(chProvider)
	if err != nil {
		return nil, nil, err
	}
	return sdk, client, nil
}
