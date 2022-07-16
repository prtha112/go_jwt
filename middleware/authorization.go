package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signature string
		var err error

		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error getting env, not comming through %v", err)
		} else {
			fmt.Println("We are getting the env values")
		}

		signature = os.Getenv("JWT_SIGNATURE")
		s := c.Request.Header.Get("Authorization")

		token := strings.TrimPrefix(s, "Bearer ")

		if err := validateToken(token, signature); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func validateToken(token string, signature string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(signature), nil
	})

	return err
}
