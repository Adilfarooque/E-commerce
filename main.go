package main

import (
	"fmt"
	"log"

	"github.com/Adilfarooque/Footgo/config"
	"github.com/Adilfarooque/Footgo/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config file")
	}
	fmt.Println(cfig)
	db, err := db.ConnectDatabase(cfig)
	if err != nil {
		log.Fatal("Error connecting to the database:%v", err)
	}
	routes := gin.Default()
	routes.LoadHTMLFiles("template/*")
	userGroup := routes.Group("/user")
	adminGroup := routes.Group("/admin")
}
