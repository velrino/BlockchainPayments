package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/blockcypher/gobcy"
)

type TransactionController struct{

}

func (TransactionController) Wallet(c *gin.Context) {

	coin := c.Param("coin")

	btc := gobcy.API{"f3d21ab19f2841bb978b3f4be2584c7c", coin, "test3"}

	resp, err := btc.GetWallet("velrino")
		
    if err == nil {
		response = resp
	}	else if err != nil  {
		response = gin.H{"message": "Not Found"}
	}
	
	c.JSON(http.StatusOK, response ) 
}