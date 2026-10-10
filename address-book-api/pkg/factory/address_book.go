package factory

import (
	"address-book-api/internal/domain/service"
	"address-book-api/internal/infrastructure/transport/controller"
	"address-book-api/internal/infrastructure/transport/route"

	"github.com/gin-gonic/gin"
)

func AddressBookFactory(router *gin.Engine) {
	service := service.AddressBookService{}
	controller := controller.NewAddressBookController(&service)
	route.SetupAddressRoutes(router, controller)
}
