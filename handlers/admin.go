package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	models "main/db"
	"main/utils"
)

type CreateAdmin struct {
	FirstName string `json:"firstName" validate:"required,min=5,max=20"`
	LastName  string `json:"lastName" validate:"required,min=5,max=20"`
	Age       int    `json:"age" validate:"gte=18,lte=120"`
	Email     string `json:"email" validate:"required,email"`
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
	validate := c.MustGet("validator").(*validator.Validate)
	if err := validate.Struct(admin); err != nil {

		err := utils.SetResponseError(http.StatusBadRequest, "400", &err)
		c.Error(err)
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

	data := utils.SetResponseData(&admin)
	c.JSON(http.StatusOK, data)
}
