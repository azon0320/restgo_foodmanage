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

func SellerLogin(ctx *gin.Context) {
	creden := &models.PasswordCredentialRequest{}
	if !util.CheckParamsValid(ctx, creden) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	resp, err := sdk.ClientExecute(models2.UnAuthLogin, [][]byte{
		[]byte(fmt.Sprint(models2.OperatorSeller)), []byte(creden.Id), []byte(creden.Password)})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func SellerCreateProduct(ctx *gin.Context) {
	input := &models.TokenCredential{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	tokenByte, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_ADDPRODUCT, [][]byte{tokenByte})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func SellerSellOnProduct(ctx *gin.Context) {
	input := &models.IdRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_TAKEONSELL, [][]byte{
		token, []byte(input.Id),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func SellerSellOffProduct(ctx *gin.Context) {
	input := &models.IdRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_TAKEOFFSELL, [][]byte{
		token, []byte(input.Id),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func SellerTransmit(ctx *gin.Context) {
	type TransmitRequest struct {
		*models.IdRequest
		TransporterId string `json:"transporter_id"`
	}
	input := &TransmitRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	resp, err := sdk.ClientExecute(models2.OPERATE_TRANSMIT, [][]byte{
		token, []byte(input.Id), []byte(input.TransporterId),
	})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}

func SellerCancelTransaction(ctx *gin.Context) {
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

func SellerUpdateProduct(ctx *gin.Context) {
	input := &models.SellerProductInfoRequest{}
	if !util.CheckParamsValid(ctx, input) {
		util.JsonReturn(ctx, models2.WithError(consts.CodeErrorParams, consts.MsgErrorParams))
		return
	}
	token, _ := json.Marshal(models2.CreateCredentialsWithToken(input.Token))
	req, _ := json.Marshal(&models2.ProductUpdateRequest{
		DataModel:            &models2.DataModel{Id: input.Id},
		EachPrice:            uint64(input.EachPrice),
		Description:          input.Description,
		Inventory:            uint32(input.Inventory),
		TransportAmount:      uint64(input.TransportAmount),
		SpecifiedTemperature: byte(input.SpecifiedTemperature),
	})
	resp, err := sdk.ClientExecute(models2.OPERATE_UPDATE_PRODUCT, [][]byte{token, req})
	if err != nil {
		util.InternalReturn(ctx)
		return
	}
	util.StringReturn(ctx, string(resp.Payload))
}
