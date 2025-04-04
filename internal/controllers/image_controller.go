package controllers

import (
	"github.com/Gabriel-Schiestl/image-processor/internal/application/usecases"
)

type imageController struct {
	Controller *controller
	processImageUseCase *usecases.ProcessImageUseCase
}

func NewImageController(controller *controller, uc *usecases.ProcessImageUseCase) *imageController {
	return &imageController{Controller: controller, processImageUseCase: uc}
}

//func (ic *imageController) RegisterRoutes() {
	//ic.Controller.Server.POST("/image", ic.ProcessImage)
//}

//func (ic *imageController) ProcessImage(c *gin.Context) {
	//file, err := c.FormFile("image")
	//if err != nil {
		//c.JSON(400, gin.H{"error": "No image provided"})
		//return
	//}

	//ic.processImageUseCase.Execute(file)
	//if er != nil {
		//c.JSON(400, gin.H{"error": err})
	//}
//}