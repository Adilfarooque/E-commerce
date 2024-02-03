package routes

import (
	"github.com/Adilfarooque/Footgo/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", handlers.UserSignUp)
	r.POST("/Userlogin",handlers.Userlogin)
	return r
}
