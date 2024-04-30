package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"main/db"

	routes "main/routes"
)

func main() {
	app := gin.Default()

	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// connect db
	db, err := db.Init()
	if err != nil {
		log.Fatalf("Error connect DB")
	}

	// set cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	app.Use(cors.New(config))

	// set routes
	api := app.Group("/api")
	routes.Set(api, db)

	println("=========================")
	println(fmt.Sprintf("APP_ENV=%s", os.Getenv("APP_ENV")))
	println(fmt.Sprintf("PORT=%s", os.Getenv("PORT")))
	println("=========================")
	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
