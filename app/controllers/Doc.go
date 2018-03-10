package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Documentation(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/doc")
}