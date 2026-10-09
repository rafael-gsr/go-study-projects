package route

import (
	"address-book-api/pkg/factory"

	"github.com/gin-gonic/gin"
)

func configAddressRoutes(router *gin.Engine) {
	controller := factory.AddressBookFactory()

	router.GET("/ping", controller.Pong)
}
