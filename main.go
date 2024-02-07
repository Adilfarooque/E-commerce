package main

import (
	"fmt"
	"log"

	"github.com/Adilfarooque/Footgo/config"
	"github.com/Adilfarooque/Footgo/db"
	"github.com/Adilfarooque/Footgo/routes"
	"github.com/gin-gonic/gin"
)

// @title Go + Gin Footgo E-Commerce API
// @description Footgo is an E-commerce platform to purchasing and selling shoes
// @contact.name API Support
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
// @query.collection.format multi

func main() {
	cfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config file")
	}
	fmt.Println(cfig)
	db, err := db.ConnectDatabase(cfig)
	if err != nil {
		log.Fatalf("Error connecting to the database:%v", err)
	}
	router := gin.Default()
	router.LoadHTMLFiles("template/*")
	userGroup := router.Group("/user")
	adminGroup := router.Group("/admin")
	routes.UserRoutes(userGroup, db)
	routes.AdminRoutes(adminGroup, db)

	listenAddres := fmt.Sprintf("%s:%s", cfig.DBPort, cfig.DBHost)
	fmt.Printf("Starting server on %s..\n", cfig.BASE_URL)
	if err := router.Run(cfig.BASE_URL); err != nil {
		log.Fatalf("Error starting server on %s:%v", listenAddres, err)
	}
}
