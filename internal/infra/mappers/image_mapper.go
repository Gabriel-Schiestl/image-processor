package mappers

import (
	"bytes"
	"image"
	"image/png"

	"github.com/Gabriel-Schiestl/image-processor/internal/domain/models"
	"github.com/Gabriel-Schiestl/image-processor/internal/infra/entities"
)

type ImageMapper struct {
}

func (m *ImageMapper) ToDomainModel(image models.Image) *entities.ImageModel {
	return &entities.ImageModel{
		Prediction: image.Prediction,
		Image:      imageToBytes(image.Image),
	}
}

func imageToBytes(img image.Image) []byte {
	var buf bytes.Buffer

	err := png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
