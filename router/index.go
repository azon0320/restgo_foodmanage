package router

import (
	"github.com/dormao/restgo_foodmanage/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.POST("/register", controllers.RegisterOperator)
	seller := router.Group("/seller")
	{
		seller.POST("/login", controllers.SellerLogin)
		seller.POST("/create_prod", controllers.SellerCreateProduct)
		seller.POST("/sell_on", controllers.SellerSellOnProduct)
		seller.POST("/sell_off", controllers.SellerSellOffProduct)
		seller.POST("/transmit", controllers.SellerTransmit)
		seller.POST("/cancel", controllers.SellerCancelTransaction)
		seller.POST("/update_prod", controllers.SellerUpdateProduct)
	}
	buyer := router.Group("/buyer")
	{
		buyer.POST("/login", controllers.BuyerLogin)
		buyer.POST("/purchase", controllers.BuyerBuyProduct)
		buyer.POST("/cancel", controllers.BuyerCancelTransaction)
		buyer.POST("/confirm", controllers.BuyerConfirmTransaction)
	}
	tspr := router.Group("/transporter")
	{
		tspr.POST("/login", controllers.TransporterLogin)
		tspr.POST("/update", controllers.TransporterUpdateTransport)
		tspr.POST("/complete", controllers.TransporterCompleteTransport)
		tspr.POST("/cancel", controllers.TransporterCancelTransport)
	}
}
