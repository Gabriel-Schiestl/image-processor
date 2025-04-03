package main

import (
	"image"

	"github.com/Gabriel-Schiestl/image-processor/internal/consumer"
	"github.com/Gabriel-Schiestl/image-processor/internal/domain/models"
)

var workers = 3

func main() {
	// server := gin.Default()
	imgCh := make(chan image.Image, 10)

	// controller := controllers.NewController(server)
	// processImgUseCase := usecases.NewProcessImageUseCase(imgCh)
	// imgController := controllers.NewImageController(controller, processImgUseCase)
	// imgController.RegisterRoutes()
	rabbitmq := models.NewRabbitMQ("images")
	defer rabbitmq.Close()

	msgs, err := rabbitmq.Consume()
	if err != nil {
		panic(err)
	}

	for i := range workers {
		go consumer.Consume(i + 1, msgs, imgCh)
	}

	for img := range imgCh {
		// Process the image here
		// For example, save it to disk or send it to another service
	}

	// server.Run(":8080")
}