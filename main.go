package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Gabriel-Schiestl/image-processor/internal/config"
	"github.com/Gabriel-Schiestl/image-processor/internal/consumer"
	"github.com/Gabriel-Schiestl/image-processor/internal/domain/models"
	"github.com/Gabriel-Schiestl/image-processor/internal/infra/repositories"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var workers = 3

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// server := gin.Default()
	imgCh := make(chan models.Image, 10)

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	dbConfig := config.NewDbConfig(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), dbPort).ToString()
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDb.Close()

	fmt.Println("Connected to database")

	repo := repositories.NewImageRepository(db)

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
		fmt.Println("Saving image:", img.Prediction)
		repo.Save(img)
	}

	// server.Run(":8080")
}