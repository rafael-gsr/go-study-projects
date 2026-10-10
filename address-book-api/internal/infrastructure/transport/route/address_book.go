package route

import (
	"address-book-api/internal/infrastructure/transport/controller"

	"github.com/gin-gonic/gin"
)

func SetupAddressRoutes(router *gin.Engine, controller *controller.AddressBookController) {
	addressRoutesRouter := router.Group("/address-book")

	addressRoutesRouter.GET("/ping", controller.Pong)
}
