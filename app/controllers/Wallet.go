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
	code := http.StatusOK

    if err == nil {
		response = resp
	}	else if err != nil  {
		code = http.StatusNotFound;		
		response = gin.H{"message": "Not Found"}
	}
	
	c.JSON(code, response ) 
}

func (WalletController) Transaction(c *gin.Context) {

	address := c.Param("address")

	btc := gobcy.API{"f3d21ab19f2841bb978b3f4be2584c7c", "btc", "test3"}

	resp, err := btc.GetAddrFull(address, nil)
	
	code := http.StatusOK

    if err == nil {
		response = resp
	}	else if err != nil  {
		code = http.StatusNotFound;		
		response = gin.H{"message": "Not Found"}
	}
	
	c.JSON(code, response ) 
}

