package route

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	configAddressRoutes(router)
}
