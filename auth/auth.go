package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Login struct {
	Username string `json:"jwt_username"`
	Password string `json:"jwt_password"`
}

func LoginHandler(c *gin.Context) {
	// implement login logic here
	var login Login
	var signature, username, password string
	var token *jwt.Token
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	signature = os.Getenv("JWT_SIGNATURE")
	username = os.Getenv("JWT_AUTH_USER")
	password = os.Getenv("JWT_AUTH_PASS")

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if username != login.Username || password != login.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Auth user not allow.",
		})
		return
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	ss, err := token.SignedString([]byte(signature))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}
