package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func productRoutes(routers *gin.RouterGroup, db *gorm.DB) {
	productRoute := routers.Group("/v1/product")
	{
		productRoute.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Product API2.",
			})
		})

	}
}
