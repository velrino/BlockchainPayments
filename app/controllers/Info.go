package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Config "github.com/velrino/BlockchainPayments/app/config"
)

func Calculate(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message":"Construct !"})
}