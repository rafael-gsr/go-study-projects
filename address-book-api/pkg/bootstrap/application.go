package bootstrap

import (
	"fmt"

	"address-book-api/internal/infrastructure/env"
	"address-book-api/pkg/route"

	"github.com/gin-gonic/gin"
)

func BootstrapApplication() {
	env.Load()

	router := gin.Default()
	route.Config(router)
	address := fmt.Sprintf(":%s", env.PORT)

	router.Run(address)
}
