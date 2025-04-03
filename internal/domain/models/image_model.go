package models

import "image"

type Image struct {
	Prediction string
	Image      image.Image
}