package consumer

import (
	"fmt"
	"mime/multipart"
)

func Consume(workerId int, ch <-chan *multipart.FileHeader) {
	for img := range ch {
		fmt.Println(img)
	}
}