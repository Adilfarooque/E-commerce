package main

import (
	"fmt"
	"log"

	"github.com/Adilfarooque/Footgo_Ecommerce/config"
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/routes"
	"github.com/gin-gonic/gin"
)

// @title Go + Gin Footgo_E-Commerce API
// @version 1.0.0
// @description Footgo is an E-commerce platform to purchasing and selling shoes
// @contact.name API Support
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
// @query.collection.format multi

func main() {
	confg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config file")
	}
	fmt.Println(confg)
	db, err := db.ConnectDatabase(confg)
	if err != nil {
		log.Fatalf("Error connecting to the database:%v", err)
	}
	r := gin.Default()
	r.LoadHTMLFiles("templates/*")
	userGroup := r.Group("/user")
	adminGroup := r.Group("/admin")
	routes.UserRoutes(userGroup, db)
	routes.AdminRoutes(adminGroup, db)
}
