package app

import (
	"github.com/gin-gonic/gin"
	. "github.com/velrino/BlockchainPayments/app/controllers"
)

func Routes()  {
	router := gin.Default()

	router.Static("/doc", "./doc")

	Wallet := router.Group("/wallet")
	{
		Wallet.GET("/:coin", TransactionController{}.Wallet)
	}

	CryptoCurrency := router.Group("/coin")
	{
		CryptoCurrency.GET("/:coin",  InfoController{}.Get)
		CryptoCurrency.GET("/:coin/:value",  InfoController{}.Calculate)
	}

	Transaction := router.Group("/transaction")
	{
		Transaction.GET("/:id",  TransactionController{}.Get)
		Transaction.POST("/",  TransactionController{}.Create)
	}

	router.Run(":80")
}