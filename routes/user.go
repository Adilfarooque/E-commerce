package routes

import (
	"github.com/Adilfarooque/Footgo/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", handlers.UserSignUp)
	r.POST("/Userlogin", handlers.Userlogin)

	products := r.Group("/products")
	{
		products.GET("", handlers.ShowAllProducts)
		products.POST("/filter",handlers.FilerCategory)
		products.GET("/image",handlers.ShowImages)
	
	}
	return r
}
