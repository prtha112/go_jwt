package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go_jwt/models"
	"go_jwt/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB *gorm.DB
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	err = server.DB.Debug().AutoMigrate(&models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	r := gin.Default()
	routes.Serve(r)
	r.Run(addr)
}
