package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Set(routers *gin.RouterGroup, db *gorm.DB) {

	adminRoutes(routers, db)

	routers.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GOLANG RESTFUL API",
		})
	})
}
