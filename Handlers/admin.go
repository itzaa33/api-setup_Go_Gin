package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	models "main/db"
)

type CreateAdmin struct {
	FirstName string `json:"firstName" binding:"required,min=5,max=20"`
	LastName  string `json:"lastName" binding:"required,min=5,max=20"`
	Age       int    `json:"age" binding:"gte=18,lte=120"`
	Email     string `json:"email" binding:"required,email"`
}

func (h *CustomerHandler) FindAll(c *gin.Context) {

	var admins []models.Admin
	err := h.Database.Find(&admins)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	}

	c.JSON(http.StatusOK, gin.H{"results": &admins})
}

func (h *CustomerHandler) Create(c *gin.Context) {

	admin := CreateAdmin{}

	if err := c.ShouldBindJSON(&admin); err != nil {
		fmt.Println("Database error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"meassage": err})

		return
	}

	result := h.Database.Create(&models.Admin{
		FirstName: admin.FirstName,
		LastName:  admin.LastName,
		Age:       admin.Age,
		Email:     &admin.Email,
	})

	if result.Error != nil {
		fmt.Println("Database error:", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"meassage": result.Error})

		return
	}

	c.JSON(http.StatusOK, gin.H{"results": &admin})
}
