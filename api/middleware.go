package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RootAuthMiddleware(c *gin.Context) {
	wallet := c.GetHeader("Wallet")
	walletroot := os.Getenv("ROOT_WALLET")
	if wallet != walletroot {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: not root"})
		c.Abort()
		return
	}
	c.Next()
}
