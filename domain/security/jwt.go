package security

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
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

func UnpackJwtToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
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

	expiration := token.Claims.(jwt.MapClaims)["exp"]

	if expiration == nil || int64(expiration.(float64)) < time.Now().Unix() {
		fmt.Println("Token is expired")
		return false, nil
	}

	fmt.Println(expiration)

	return token.Valid, err
}

func GetTokenString(s string) (string, error) {
	token := strings.Split(s, " ")[1]
	if token == "" {
		return "", errors.New("token is empty")
	}
	return token, nil
}

func ExtractClaimFromHeader(tokenString string, target string) (string, error) {

	token, err := UnpackJwtToken(tokenString, os.Getenv("SECRET"))

	if err != nil {
		return "", err
	}

	return ExtractClaims(token)[target].(string), nil

}
