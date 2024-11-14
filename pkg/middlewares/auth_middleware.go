package middlewares

import (
	"fmt"
	"os"
	"strings"
	"time"
	"user_api/pkg/utils/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func AuthMiddleware(c *gin.Context) {
	authToken, authErr := getToken(c.GetHeader("Authorization"))
	if authErr != nil {
		c.JSON(authErr.Status, authErr)
		c.Abort()
		return
	}

	claims, err := parseToken(authToken)
	if err != nil {
		restErr := errors.NewUnauthorizedError(err.Error())
		c.JSON(restErr.Status, restErr)
		c.Abort()
		return
	}

	if claims.UserID != "" && claims.Email != "" {
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
	} else {
		restErr := errors.NewUnauthorizedError("Invalid token.")
		c.JSON(restErr.Status, restErr)
		c.Abort()
		return
	}
	c.Next()
}

func getToken(authHeader string) (string, *errors.RestErr) {
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		return bearerToken[1], nil
	}
	return "", errors.NewUnauthorizedError("Invalid token format.")
}

func parseToken(authToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(authToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("token expired")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
