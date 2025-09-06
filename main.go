package main

import (
	"github.com/asentientbanana/pausalkole-admin/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := common.InitDatabase()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	common.InitializeRoutes(server, db)

	err = server.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		panic(err)
	}
}
