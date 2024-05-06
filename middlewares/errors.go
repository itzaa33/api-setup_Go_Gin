package middleware

import (
	"net/http"

	"main/utils"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case utils.ResponseError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		}
	}
}
