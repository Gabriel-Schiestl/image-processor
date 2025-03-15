package usecases

import (
	"errors"
	"mime/multipart"
	"path/filepath"
)

var extensionsAllowed = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".bmp": true, ".webp": true}

type ProcessImageUseCase struct {
	imagesCh chan<- *multipart.FileHeader
}

func NewProcessImageUseCase(ch chan<- *multipart.FileHeader) *ProcessImageUseCase {
	return &ProcessImageUseCase{imagesCh: ch}
}

func (uc *ProcessImageUseCase) Execute(file *multipart.FileHeader) error {
	if ext := filepath.Ext(file.Filename); !extensionsAllowed[ext] {
		return errors.New("not allowed extension")
	}

	uc.imagesCh <- file

	return nil
}