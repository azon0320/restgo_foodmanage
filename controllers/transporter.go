package controllers

import (
	"encoding/json"
	"fmt"
	models2 "github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/models/consts"
	"github.com/dormao/restgo_foodmanage/models"
	"github.com/dormao/restgo_foodmanage/sdk"
	"github.com/dormao/restgo_foodmanage/util"
	"github.com/gin-gonic/gin"
)

func TransporterLogin(ctx *gin.Context) {
	creden := &models.PasswordCredentialRequest{}
	if !util.CheckParamsValid(ctx, creden) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	resp, err := sdk.ClientExecute(models2.UnAuthLogin, [][]byte{
		[]byte(fmt.Sprint(models2.OperatorTransporter)), []byte(creden.Id), []byte(creden.Password)})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func TransporterCancelTransport(ctx *gin.Context) {
	input := &models.IdRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_CANCELTRANSPORT, [][]byte{
		token, []byte(input.Id),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func TransporterCompleteTransport(ctx *gin.Context) {
	input := &models.IdRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_COMPLETE_TRANSPORT, [][]byte{
		token, []byte(input.Id),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func TransporterUpdateTransport(ctx *gin.Context) {
	input := models.TransporterUpdateDetailsRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	req, _ := json.Marshal(&models2.TransportDetails{Temperature: byte(input.Temperature)})
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_UPDATE_TRANSPORT, [][]byte{
		token, []byte(input.Id), req,
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}
