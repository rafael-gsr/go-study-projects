package bootstrap

import (
	"fmt"

	"address-book-api/internal/core/middleware"
	"address-book-api/internal/infrastructure/env"
	"address-book-api/pkg/factory"

	"github.com/gin-gonic/gin"
)

func BootstrapApplication() {
	env.Load()

	router := gin.Default()

	middleware.SetupMiddlewares(router)
	address := fmt.Sprintf(":%s", env.PORT)

	factory.AddressBookFactory(router)

	router.Run(address)
}
