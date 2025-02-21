package main

import (
	"log"
	"my-go-project/internal/database"
	"my-go-project/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Thiết lập chế độ release cho Gin
	gin.SetMode(gin.ReleaseMode)

	// Kết nối với database
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Không thể kết nối tới database: %v", err)
	}

	// Thực hiện migration các model
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Migration thất bại: %v", err)
	}

	// Khởi tạo router và đăng ký các route
	router := gin.Default()
	routes.RegisterRoutes(router, db)

	// Chạy server trên cổng 8080
	router.Run(":8080")
}