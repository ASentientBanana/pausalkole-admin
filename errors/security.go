package errors

import "github.com/gin-gonic/gin"

func CreateTokenInValidError() gin.H {

	return gin.H{"error": "Token is invalid"}
}
