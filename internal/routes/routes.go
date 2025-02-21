package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-project/internal/controllers"
)

// RegisterRoutes đăng ký các route và liên kết với controller
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Route cho country
	router.POST("/add-country", controllers.AddCountry(db))
	router.GET("/get-all-countries", controllers.GetAllCountries(db))
}
