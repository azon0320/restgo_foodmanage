package models

import "github.com/dormao/chaincode_foodmanage/models"

type DataProductSnapshot struct {
	*models.ProductSnapshot
}

type DataTransactionSnapshot struct {
	*models.TransactionOrder
	*models.TransportOrder
}

type DataUser struct {
	*models.DataModel
	Description string
}
