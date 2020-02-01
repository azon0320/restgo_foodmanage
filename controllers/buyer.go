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

func BuyerLogin(ctx *gin.Context) {
	creden := &models.PasswordCredentialRequest{}
	if !util.CheckParamsValid(ctx, creden) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	resp, err := sdk.ClientExecute(models2.UnAuthLogin, [][]byte{
		[]byte(fmt.Sprint(models2.OperatorBuyer)), []byte(creden.Id), []byte(creden.Password)})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func BuyerBuyProduct(ctx *gin.Context) {
	type PurchaseRequest struct {
		*models.IdRequest
		BuyCount int `json:"count"`
	}
	input := &PurchaseRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_PURCHASE, [][]byte{
		token, []byte(input.Id), []byte(fmt.Sprint(input.BuyCount)),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func BuyerConfirmTransaction(ctx *gin.Context) {
	input := &models.IdRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_CONFIRM, [][]byte{
		token, []byte(input.Id),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func BuyerCancelTransaction(ctx *gin.Context) {
	input := &models.IdRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_CANCELORDER, [][]byte{
		token, []byte(input.Id),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}
