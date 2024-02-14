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
	gin.SetMode(gin.ReleaseMode)
	//Load application configuration
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config file")
	}
	//Connect to the database
	db, err := db.ConnectDatabase(conf)
	if err != nil {
		log.Fatalf("Error connecting to the database:%v", err)
	}
	//Initial Gin router
	r := gin.Default()
	//Load Html templates
	//r.LoadHTMLFiles("template/*")
	//Define route groups
	userGroup := r.Group("/user")
	adminGroup := r.Group("/admin")
	//Register the routes
	routes.UserRoutes(userGroup, db)
	routes.AdminRoutes(adminGroup, db)

	listenAddress := fmt.Sprintf("%s:%s", conf.DBPort, conf.DBHost)
	fmt.Printf("Starting sever on %s..\n", conf.BASE_URL)
	if err := r.Run(conf.BASE_URL); err != nil {
		log.Fatalf("Error starting server on %s:%s", listenAddress, err)
	}
}
