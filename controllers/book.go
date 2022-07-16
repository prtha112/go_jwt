package controllers

import (
	"go_jwt/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct{}

func (server *Server) FindAll(c *gin.Context) {
	var err error
	var books []models.Book
	err = server.DB.Debug().Model(&models.Book{}).Limit(100).Find(&books).Error
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &books,
	})

}
