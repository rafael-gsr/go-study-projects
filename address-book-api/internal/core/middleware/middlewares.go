package middleware

import "github.com/gin-gonic/gin"

func SetupMiddlewares(r *gin.Engine) {
	r.Use(panicHandler())
}
