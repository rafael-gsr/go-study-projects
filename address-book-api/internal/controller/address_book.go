// Package controller contains the application controllers to connect routes to services
package controller

import (
	"address-book-api/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type AddressBookController struct {
	service service.AddressBookService
}

func (c *AddressBookController) SetService(service service.AddressBookService) {
	c.service = service
}

func (c *AddressBookController) Pong(ginContext *gin.Context) {
	serviceResponse := c.service.Pong()

	ginContext.JSON(200, serviceResponse)
}
