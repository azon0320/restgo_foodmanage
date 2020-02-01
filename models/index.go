package models

type TokenCredential struct {
	Token string `json:"token" form:"token" binding:"required"`
}
