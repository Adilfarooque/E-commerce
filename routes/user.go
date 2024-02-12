package routes

import (
	"github.com/Adilfarooque/Footgo/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", handlers.UserSignUp)
	r.POST("/login", handlers.Userlogin)

	r.POST("/send-otp", handlers.SendOtp)
	r.POST("/verify-otp", handlers.VerifyOtp)

	r.POST("/forgot-password", handlers.ForgotPasswordSend)

	products := r.Group("/products")
	{
	products.GET("", handlers.ShowAllProducts)
	products.POST("/filter", handlers.FilerCategory)
	products.GET("/image", handlers.ShowImages)
	}
	return r

}
