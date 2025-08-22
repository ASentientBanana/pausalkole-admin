package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJwtToken(id string) (string, error) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		panic("secret env variable not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiry
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(secret))

}

func UnpackJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})
	return token, err
}

func ExtractClaims(token *jwt.Token) map[string]interface{} {
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

func ValidateJwtToken(tokenString string, secret string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure signing method is correct
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return []byte(secret), nil
	})

	experation := token.Claims.(jwt.MapClaims)["exp"]

	if experation == nil || experation.(int64) < time.Now().Unix() {
		fmt.Println("Token is expired")
		return false, nil
	}

	fmt.Println(experation)

	return token.Valid, err
}
