package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	. "github.com/velrino/BlockchainPayments/app/controllers"
)

func Routes()  {
	router := gin.Default()

	router.Use(cors.Default())

	router.Static("/doc", "./doc")
	
	CryptoCurrency := router.Group("/coin")
	{
		CryptoCurrency.GET("/:coin",  CryptoController{}.Get)
		CryptoCurrency.GET("/:coin/:value",  CryptoController{}.Calculate)
	}

	Transaction := router.Group("/transaction")
	{
		Transaction.GET("/:id",  TransactionController{}.Get)
		Transaction.POST("/",  TransactionController{}.Create)
	}

	Wallet := router.Group("/wallet")
	{
		Wallet.GET("/", WalletController{}.Address)
		Wallet.GET("/:address", WalletController{}.Transaction)
	}

	router.Run(":80")
}