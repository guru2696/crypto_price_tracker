package main

import (
	"crypto_price_tracker/plog"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/crypto/price", GetCryptoPriceHandler)
	plog.Info("Starting Server on 8080", nil)

	r.Run(":8080")
}
