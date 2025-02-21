package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
            c.Abort()
            return
        }
        
        // Thêm logic kiểm tra token ở đây
        
        c.Next()
    }
} 