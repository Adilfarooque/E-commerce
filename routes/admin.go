package routes

import (
	"github.com/Adilfarooque/Footgo/handlers"
	"github.com/Adilfarooque/Footgo/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/adminlogin", handlers.LoginHandler)

	r.Use(middleware.AdminAuthMiddleware())
	{
		r.GET("/dashboard", handlers.DashBoard)
	}

	users := r.Group("/users")
	{
		users.GET("", handlers.GetUsers)
		users.PUT("/block", handlers.BlockUser)
		users.PUT("/unblock", handlers.UnBlockUser)
	}
	return r
}
