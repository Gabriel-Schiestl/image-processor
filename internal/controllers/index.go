package controllers

import "github.com/gin-gonic/gin"

type controller struct {
	Server *gin.Engine
}

func NewController(server *gin.Engine) *controller {
	return &controller{Server: server}
}