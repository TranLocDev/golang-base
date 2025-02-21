package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-project/internal/models"
)

// AddCountry nhận payload và thêm một bản ghi country mới vào database.
func AddCountry(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCountry models.Countries

		if err := c.BindJSON(&newCountry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newCountry).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Country Added Successfully",
			"newCountry": newCountry,
		})
	}
}

// GetAllCountries trả về danh sách tất cả country trong database.
func GetAllCountries(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var allCountries []models.Countries

		if err := db.Find(&allCountries).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "allCountries": allCountries})
	}
}
