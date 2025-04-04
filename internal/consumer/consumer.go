package consumer

import (
	"github.com/Gabriel-Schiestl/image-processor/internal/application/usecases"
	"github.com/Gabriel-Schiestl/image-processor/internal/domain/models"
	"github.com/rabbitmq/amqp091-go"
)

func Consume(workerId int, ch <-chan amqp091.Delivery, imgCh chan<- models.Image) {
	useCase := usecases.NewProcessImageUseCase(imgCh)

	for img := range ch {
		go func(img amqp091.Delivery) {
			useCase.Execute(img.Body)
		}(img)
	}
}