package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"strconv"
	. "github.com/velrino/BlockchainPayments/app/responses"
)

type CryptoController struct{}

var (
	baseURL  = "https://api.coinmarketcap.com/v1"
	tickerBRL  = "https://api.coinmarketcap.com/v1/ticker/COIN/?convert=BRL"
)

func (CryptoController) Get(c *gin.Context) {

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

func (CryptoController) Calculate(c *gin.Context) {

	coin := c.Param("coin")
	value, _ := strconv.ParseFloat(c.Param("value") ,64)
	
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
	
	PriceBrl, _ := strconv.ParseFloat(data[0].PriceBrl ,64)
		
	data[0].Calculate = (value/PriceBrl)
	
	c.JSON(http.StatusOK, data[0] ) 
	return
}

