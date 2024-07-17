package middleware

import (
	"net/http"

	tokenhelper "github.com/Aman123at/gin-userauth/helper"
	"github.com/gin-gonic/gin"
)

func IsUserLoggedIn(c *gin.Context) {
	cToken := c.Request.Header.Get("Authorization")

	if cToken == "" {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "No authorization token found"})
		c.Abort()
		return
	}

	message := tokenhelper.ValidateToken(cToken)

	if message != "valid" {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": message})
		c.Abort()
		return
	}

	c.Next()
}
