package interfaces

import "github.com/Gabriel-Schiestl/image-processor/internal/domain/models"

type ImageRepository interface {
	Save(image models.Image) error
}