package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/blockcypher/gobcy"
	. "github.com/velrino/BlockchainPayments/app/models"
	. "github.com/velrino/BlockchainPayments/app/dao"
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

var dao = DataBaseParams{}

func (TransactionController) List(c *gin.Context) {
	var model TransactionModel
	Response, err := dao.FindById("transacions",model,"5aa5cc722f2eeb496f8c2fc8")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Response})
}

func (TransactionController) CreateTransaction(c *gin.Context) {
	var model TransactionModel
	if err := c.ShouldBindJSON(&model); err == nil {	
		model.ID = bson.NewObjectId();
		model.Status = "waiting";
		model.CreatedAt = time.Now();

		Response, err := dao.Insert("transacions",model)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": Response})
		
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}