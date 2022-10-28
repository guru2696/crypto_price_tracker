package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetCryptoPriceHandler(c *gin.Context) {
	redisClient := GetRedisClient()
	defer redisClient.Close()

	cache := CheckRedisForPrice(redisClient)
	if cache != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cache,
		})
		return
	}

	resp, err := GetBitCoinPrice()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot Decode JSON",
		})
	}
	formattedData := map[string]map[string]string{
		"bitcoin": {
			"EUR": resp.Bpi["EUR"].Rate,
			"USD": resp.Bpi["USD"].Rate,
		},
	}

	expirySeconds, _ := strconv.Atoi(os.Getenv("EXPIRY_SECONDS"))
	fmt.Println(expirySeconds)
	redisClient.Set(c, "bitcoin_price_USD", formattedData["bitcoin"]["USD"], time.Duration(expirySeconds)*time.Second)
	redisClient.Set(c, "bitcoin_price_EUR", formattedData["bitcoin"]["EUR"], time.Duration(expirySeconds)*time.Second)

	c.JSON(http.StatusOK, gin.H{
		"data": formattedData,
	})
}
