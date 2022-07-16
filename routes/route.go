package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go_jwt/auth"
	"go_jwt/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Serve(r *gin.Engine) {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	Authorization := middleware.AuthorizationMiddleware()

	signature := os.Getenv("JWT_SIGNATURE")
	tokenTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_TIME"))
	Apikey := os.Getenv("API_KEY")

	loginConfig := auth.Env{
		Jwtsignature:  signature,
		Jwtexpiretime: tokenTime,
		Apikey:        Apikey,
	}

	r.POST("/login", func(c *gin.Context) {
		auth.LoginHandler(c, loginConfig)
	})

	protected := r.Group("/", Authorization)
	protected.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	protected.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
