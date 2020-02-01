package controllers

import (
	ccmodels "github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/models/consts"
	"github.com/dormao/restgo_foodmanage/models"
	"github.com/dormao/restgo_foodmanage/sdk"
	"github.com/dormao/restgo_foodmanage/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func RegisterOperator(ctx *gin.Context) {
	request := &models.RegisterRequest{}
	if ctx.ShouldBindWith(request, binding.Form) != nil {
		util.JsonReturn(ctx, ccmodels.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	args := [][]byte{[]byte(request.Password)}
	var resp channel.Response
	var err error
	switch request.OperatorType {
	case ccmodels.OperatorSeller:
		resp, err = sdk.ClientExecute(ccmodels.UnAuthRegisterSeller, args)
	case ccmodels.OperatorBuyer:
		resp, err = sdk.ClientExecute(ccmodels.UnAuthRegisterBuyer, args)
	case ccmodels.OperatorTransporter:
		resp, err = sdk.ClientExecute(ccmodels.UnAuthRegisterTransporter, args)
	default:
		util.JsonReturn(ctx, ccmodels.WithError(consts.CodeErrorOperatorNotFound, consts.MsgErrorOperatorNotFound))
		return
	}
	if err != nil {
		util.JsonReturn(ctx, ccmodels.WithError(consts.CodeErrorGeneral, consts.MsgErrorGeneral))
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

/* internal use
func Ping(ctx *gin.Context){
	err := sdk.ExecutePing()
	if err != nil {
		util.JsonReturn(ctx, ccmodels.WithError(consts.CodeErrorGeneral, consts.MsgErrorGeneral))
		return
	}
	util.StringReturn(ctx, "pong")
}

func QueryPing(ctx *gin.Context){
	resp, err := sdk.QueryPing()
	if err != nil {
		util.JsonReturn(ctx, ccmodels.WithError(consts.CodeErrorGeneral, consts.MsgErrorGeneral))
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}
*/
