package factory

import (
	"address-book-api/internal/controller"
	"address-book-api/internal/domain/service"
)

func AddressBookFactory() controller.AddressBookController {
	service := service.AddressBookService{}
	controller := controller.AddressBookController{}

	controller.SetService(service)

	return controller
}
