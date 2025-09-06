package main

import (
	"github.com/asentientbanana/pausalkole-admin/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
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

	err := server.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		panic(err)
	}
}
