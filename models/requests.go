package models

type PasswordCredentialRequest struct {
	Id       string `form:"id" binding:"required"`
	Password string `form:"password" binding:"required,min=6,max=24"`
}

type RegisterRequest struct {
	//TODO The Operator Number Type
	OperatorType int    `form:"type" binding:"min=1,max=3"`
	Password     string `form:"password" binding:"required,min=6,max=24"`
}

type IdRequest struct {
	*TokenCredential
	Id string `form:"id" binding:"required"`
}

type BuyerViewProductRequest struct {
	*TokenCredential
	Limit int `form:"limit" binding:"omitempty,min=5"`
	Page  int `form:"page" binding:"omitempty,min=1"`
}

type SellerProductInfoRequest struct {
	*TokenCredential
	*IdRequest
	EachPrice            int    `form:"price" binding:"required,min=0"`
	Description          string `form:"description" binding:"required,min=0"`
	Inventory            int    `form:"inventory" binding:"required,min=0"`
	TransportAmount      int    `form:"trans_amount" binding:"required,min=0"`
	SpecifiedTemperature int    `form:"spec_temp" binding:"required,min=-127,max=127"`
}

type TransporterUpdateDetailsRequest struct {
	*IdRequest
	Temperature int `form:"temp" binding:"required,min=-127,max=127"`
}
