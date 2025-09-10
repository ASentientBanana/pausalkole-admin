package errors

import "github.com/gin-gonic/gin"

func CreateTokenInValidError() gin.H {
	return gin.H{"error": "Token is invalid"}
}
func CreateInternalServerError() gin.H {
	return gin.H{"error": "Internal Server Error"}
}
