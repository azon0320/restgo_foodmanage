package util

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/models/consts"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
)

func CheckParamsValid(ctx *gin.Context, inputModel interface{}) bool {
	return ctx.ShouldBindWith(inputModel, binding.Form) == nil
}

func ParseDataReturns(response channel.Response, err error) (*models.DataReturns, error) {
	if err != nil {
		return nil, err
	}
	var returns *models.DataReturns
	err = json.Unmarshal(response.Payload, returns)
	if err != nil {
		return nil, err
	}
	return returns, nil
}

func JsonReturn(ctx *gin.Context, data *models.DataReturns) {
	ctx.JSON(http.StatusOK, data)
}

func InternalReturn(ctx *gin.Context) {
	JsonReturn(ctx, models.WithError(consts.CodeErrorGeneral, consts.MsgErrorGeneral))
}

func StringReturn(ctx *gin.Context, s string) {
	ctx.String(http.StatusOK, s)
}

func StringArray2ByteArray(pieces []string) [][]byte {
	argsBytes := make([][]byte, len(pieces))
	for i := 0; i < len(argsBytes); i++ {
		argsBytes[i] = []byte(pieces[i])
	}
	return argsBytes
}
