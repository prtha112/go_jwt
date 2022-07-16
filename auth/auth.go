package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Jwtusername string `json:"jwt_username"`
	Jwtpassword string `json:"jwt_password"`
}

type Env struct {
	Jwtsignature  string
	Jwtexpiretime int
	Apikey        string
}

func LoginHandler(c *gin.Context, env Env) {
	// implement login logic here
	var Apikey string
	var ExpiresAt time.Time
	var login Login
	var token *jwt.Token

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Apikey = c.Request.Header.Get("X-API-Key")
	if Apikey != env.Apikey {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Auth user not allow.",
		})
		return
	}

	ExpiresAt = time.Now().Add(5 * time.Minute)
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: ExpiresAt.Unix(),
	})

	tokenString, err := token.SignedString([]byte(env.Jwtsignature))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token":     tokenString,
		"username":  login.Jwtusername,
		"expire_at": ExpiresAt,
	})
}
