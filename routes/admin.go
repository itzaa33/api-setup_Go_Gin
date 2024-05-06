package routes

import (
	handler "main/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func adminRoutes(routers *gin.RouterGroup, db *gorm.DB) {
	ctrls := handler.CustomerHandler{Database: db}

	userRouter := routers.Group("/v1/admin")
	{
		userRouter.GET("/", ctrls.FindAll)
		userRouter.POST("/", ctrls.Create)
	}

}
