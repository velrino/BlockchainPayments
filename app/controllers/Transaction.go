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

var dao = DataBaseParams{}

type TransactionController struct{}

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

func (TransactionController) Get(c *gin.Context) {

	id := c.Param("id")

	var model TransactionModel
	idCheck := bson.IsObjectIdHex(id)
	if idCheck != true {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID invalid"})
		return
	}
	Response, err := dao.FindById("transacions",model,id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Response})
}

func (TransactionController) Create(c *gin.Context) {
	var model TransactionModel
	
	if err := c.ShouldBindJSON(&model); err == nil {	
		
		model.ID = bson.NewObjectId();
		model.Status = "waiting";
		model.CreatedAt = time.Now();

		btc := gobcy.API{"f3d21ab19f2841bb978b3f4be2584c7c", "btc", "test3"}
		_, addr, err := btc.GenAddrWallet("velrino")

		model.Blockchain = addr

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