package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/blockcypher/gobcy"
)

type WalletController struct{}

func (WalletController) Address(c *gin.Context) {

	btc := gobcy.API{"f3d21ab19f2841bb978b3f4be2584c7c", "btc", "test3"}

	resp, err := btc.GetWallet("velrino")
		
    if err == nil {
		response = resp
	}	else if err != nil  {
		response = gin.H{"message": "Not Found"}
	}
	
	c.JSON(http.StatusOK, response ) 
}

func (WalletController) Transaction(c *gin.Context) {

	address := c.Param("address")

	btc := gobcy.API{"f3d21ab19f2841bb978b3f4be2584c7c", "btc", "test3"}

	resp, err := btc.GetAddrFull(address, nil)
	
    if err == nil {
		response = resp
	}	else if err != nil  {
		response = gin.H{"message": "Not Found"}
	}
	
	c.JSON(http.StatusOK, response ) 
}

