package middleware

import (
	"fmt"
	"github.com/asentientbanana/pausalkole-admin/domain/security"
	"github.com/asentientbanana/pausalkole-admin/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("SECRET")

		if secret == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errors.CreateInternalServerError())
			c.Abort()
			return
		}

		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, errors.CreateTokenInValidError())
			c.Abort()
			return
		}

		splitHeader := strings.Split(header, " ")

		prefix := splitHeader[0]
		token := splitHeader[1]

		if prefix != "Bearer" || token == "" {
			c.JSON(401, errors.CreateTokenInValidError())
			c.Abort()
			return
		}

		isValid, err := security.ValidateJwtToken(token, secret)
		fmt.Println(header)
		if err != nil || !isValid {
			c.JSON(401, errors.CreateTokenInValidError())
			c.Abort()
			return
		}

		c.Next()
	}
}
