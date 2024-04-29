package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"main/db"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// connect db
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connect DB")
	}
	db.AutoMigrate()

	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	app.Use(cors.New(config))

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	println("=========================")
	println(fmt.Sprintf("APP_ENV=%s", os.Getenv("APP_ENV")))
	println(fmt.Sprintf("PORT=%s", os.Getenv("PORT")))
	println("=========================")
	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
