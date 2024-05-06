package main

import (
	"main/db"

	middleware "main/middlewares"
	routes "main/routes"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	app := gin.Default()
	app.Use(
		middleware.ErrorHandler(),
	)

	// log request
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Initialize the validator
	validate := validator.New()

	// Middleware to attach validator to Gin context
	app.Use(func(c *gin.Context) {
		c.Set("validator", validate)
		c.Next()
	})

	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file")
	}

	// connect db
	db, err := db.Init()
	if err != nil {
		log.Panic("Error connect DB")
	}

	// set cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	app.Use(cors.New(config))

	// set routes
	api := app.Group("/api")
	routes.Set(api, db)

	fmt.Println("=========================")
	println(fmt.Sprintf("APP_ENV=%s", os.Getenv("APP_ENV")))
	println(fmt.Sprintf("PORT=%s", os.Getenv("PORT")))
	fmt.Println("=========================")
	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
