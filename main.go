package main

import (
	"ingestion_api/configs"
	"ingestion_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.MongoConnect()

	routes.UserRoute(router)
	routes.PingRoutes(router)

	router.Run("192.168.29.89:8080")
}
