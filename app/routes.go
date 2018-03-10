package app

import (
	"github.com/gin-gonic/gin"
	Controllers "github.com/velrino/BlockchainPayments/app/controllers"
)

func Routes() {
	router := gin.Default()

	router.GET("/", Controllers.Documentation)
	router.Static("/doc", "./doc")

	API := router.Group("/api")
	API.GET("/calculate",  Controllers.Calculate)
	API.GET("/calculate/:crypto",  Controllers.Calculate)

	router.Run(":80")
}