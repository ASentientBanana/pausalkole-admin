package main

import (
	"github.com/asentientbanana/pausalkole-admin/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDatabase()

	server := gin.Default()

	server.Use(cors.Default())

	common.InitializeRoutes(server, db)

	err := server.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		panic(err)
	}
}
