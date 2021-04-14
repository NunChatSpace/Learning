package Interceptor

import (
	"fmt"
	"main/model"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "ThisIsVerySecretKey"

func JwtVerify(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(claims)
		c.Set("jwt_username", claims["username"])
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"result": "NG", "error": err})
		c.Abort()
	}
}

func JwtSign(payload model.User) (string, error) {

	authClaims := jwt.MapClaims{}
	authClaims["id"] = payload.ID
	authClaims["username"] = payload.Username
	authClaims["level"] = payload.Level
	authClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	token, err := at.SignedString([]byte(secretKey))

	return token, err
}
