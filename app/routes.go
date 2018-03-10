package app

import (
	"time"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	Controllers "github.com/velrino/BlockchainPayments/app/controllers"
)

func Routes() {
	router := gin.Default()

	router.GET("/", Controllers.Doc)
	router.Static("/doc", "./doc")

	API := router.Group("/api")
	API.GET("/calculate",  Controllers.Calculate)
	API.GET("/calculate/:crypto",  Controllers.Calculate)

	router.Run(":7878")
}