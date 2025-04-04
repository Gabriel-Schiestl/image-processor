package usecases

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"

	"github.com/Gabriel-Schiestl/image-processor/internal/domain/models"
	"github.com/nfnt/resize"
)

var extensionsAllowed = map[string]bool{"jpg": true, "jpeg": true, "png": true, "bmp": true, "webp": true}

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

	imgBuffer, err := base64.StdEncoding.DecodeString(model.Data.ImgBase64)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}
	reader := bytes.NewReader(imgBuffer)
	
	img, str, err := image.Decode(reader)
	if err != nil {
		return
	}
	fmt.Println("Image decoded:", str)
	if _, ok := extensionsAllowed[str]; !ok {
		fmt.Println("Image extension not allowed:", str)
		return
	}
	fmt.Println("Image decoded:", str)
	img = resize.Resize(224, 224, img, resize.Lanczos3)
	fmt.Println("Image resized:", img.Bounds())
	newImage := models.Image{
		Prediction: model.Data.Prediction,
		Image: 	 img,
	}
	fmt.Println("Image decoded:", newImage.Prediction)
	uc.imagesCh <- newImage
}