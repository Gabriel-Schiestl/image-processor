package controllers

type imageController struct {
	Controller *controller
}

func NewImageController(controller *controller) *imageController {
	return &imageController{Controller: controller}
}

func (ic *imageController) RegisterRoutes() {
	ic.Controller.Server.GET("/image", ic.GetImage)
}