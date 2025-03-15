package main

import (
	"mime/multipart"

	"github.com/Gabriel-Schiestl/image-processor/internal/application/usecases"
	"github.com/Gabriel-Schiestl/image-processor/internal/consumer"
	"github.com/Gabriel-Schiestl/image-processor/internal/controllers"
	"github.com/gin-gonic/gin"
)

var workers = 3

func main() {
	server := gin.Default()

	imgCh := make(chan *multipart.FileHeader, 10)

	controller := controllers.NewController(server)
	processImgUseCase := usecases.NewProcessImageUseCase(imgCh)
	imgController := controllers.NewImageController(controller, processImgUseCase)
	imgController.RegisterRoutes()

	for i := range workers {
		go consumer.Consume(i + 1, imgCh)
	}

	server.Run(":8080")
}