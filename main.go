package main

import (
	"github.com/Gabriel-Schiestl/image-processor/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	controller := controllers.NewController(server)

	server.Run(":8080")

}