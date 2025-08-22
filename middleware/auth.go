package middleware

import (
	"fmt"
	"github.com/asentientbanana/pausalkole-admin/domain/security"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("SECRET")
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, gin.H{"error": "token is empty"})
			c.Abort()
			return
		}

		splitHeader := strings.Split(header, " ")

		prefix := splitHeader[0]
		token := splitHeader[1]

		if prefix != "Bearer" || token == "" {
			c.JSON(401, gin.H{"error": "Token invalid"})
			c.Abort()
			return
		}

		isValid, err := security.ValidateJwtToken(token, secret)
		fmt.Println(header)
		if err != nil || !isValid {
			c.JSON(401, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
