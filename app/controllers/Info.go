package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Calculate(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message":"Construct !"})
}