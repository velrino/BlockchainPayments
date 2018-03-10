package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"strconv"
	. "github.com/velrino/BlockchainPayments/app/responses"
)

type InfoController struct{}

var (
	baseURL  = "https://api.coinmarketcap.com/v1"
	tickerBRL  = "https://api.coinmarketcap.com/v1/ticker/COIN/?convert=BRL"
)

func (InfoController) Get(c *gin.Context) {

	coin := c.Param("coin")
	
	var replacer = strings.NewReplacer("COIN", coin)
    str := replacer.Replace(tickerBRL)

	resp, err := http.Get(str)
	
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var data []CoinResponse
	
	json.NewDecoder(resp.Body).Decode(&data)

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        response = data[0]
    } else {
        response = gin.H{"message": "Not Found"}
	}
	
	c.JSON(resp.StatusCode, response) 
	return
}

func (InfoController) Calculate(c *gin.Context) {

	coin := c.Param("coin")
	value, err := strconv.ParseFloat(c.Param("value") ,64)
	
	if err != nil {
		response = gin.H{"message": "Fail Float"}
	}
	
	var replacer = strings.NewReplacer("COIN", coin)
	str := "https://api.coinmarketcap.com/v1/ticker/COIN/?convert=BRL"
    str = replacer.Replace(str)

	resp, err := http.Get(str)
	
	if err != nil {
		return
	}
	
	defer resp.Body.Close()
	var data []CoinResponse

	json.NewDecoder(resp.Body).Decode(&data)
	
	PriceBrl, err := strconv.ParseFloat(data[0].PriceBrl ,64)
	
	if err != nil {
		response = gin.H{"message": "Fail Float"}
	}
	
	data[0].Calculate = (value/PriceBrl)
	
	c.JSON(http.StatusOK, data[0]  ) 
	return
}

