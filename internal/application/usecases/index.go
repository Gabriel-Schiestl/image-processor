package usecases

import (
	"bytes"
	"encoding/json"
	"image"

	"github.com/Gabriel-Schiestl/image-processor/internal/domain/models"
	"github.com/nfnt/resize"
)

var extensionsAllowed = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".bmp": true, ".webp": true}

type ProcessImageUseCase struct {
	imagesCh chan<- models.Image
}

func NewProcessImageUseCase(ch chan<- models.Image) *ProcessImageUseCase {
	return &ProcessImageUseCase{imagesCh: ch}
}

func (uc *ProcessImageUseCase) Execute(msg []byte) {
	model := &models.Message{}

	if err := json.Unmarshal(msg, &model); err != nil {
		return
	}

	reader := bytes.NewReader(model.ImgBuffer)

	img, str, err := image.Decode(reader)
	if err != nil {
		return
	}

	if _, ok := extensionsAllowed[str]; !ok {
		return
	}

	img = resize.Resize(224, 224, img, resize.Lanczos3)

	newImage := models.Image{
		Prediction: model.Prediction,
		Image: 	 img,
	}

	uc.imagesCh <- newImage
}