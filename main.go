package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/crypto/price", GetCryptoPrice)

	r.Run(":8080")
}
