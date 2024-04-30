package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"main/handlers"
)

func adminRoutes(routers *gin.RouterGroup, db *gorm.DB) {
	ctrls := handlers.CustomerHandler{Database: db}

	userRouter := routers.Group("/v1/admin")
	{
		userRouter.GET("/", ctrls.FindAll)
		userRouter.POST("/", ctrls.Create)
	}
}
