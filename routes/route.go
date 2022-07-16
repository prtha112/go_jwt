package routes

import (
	"net/http"

	"go_jwt/auth"
	"go_jwt/middleware"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	// authenticate := middleware.Authenticate().MiddlewareFunc()
	// authorize := middleware.Authorize()
	Authorization := middleware.AuthorizationMiddleware()
	// r.Use(Authorization)

	r.POST("/login", auth.LoginHandler)

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
