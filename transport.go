package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCryptoPriceHandler(c *gin.Context) {
	// Check in Cache
	if cache := GetBitCoinPriceCache(); cache != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cache,
		})
		return
	}

	// Fetch from Client
	resp, err := GetBitCoinPrice()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
