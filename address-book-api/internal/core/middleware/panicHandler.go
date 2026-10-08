// Package middleware contains functins responsable for manipulate requisitions or responses before reaching the target
package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func panicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer sendInternalServerErrorMessage(c)
		c.Next()
	}
}

func sendInternalServerErrorMessage(c *gin.Context) {
	if err := recover(); err != nil {
		fmt.Printf("err: unhandled error - %s", err)

		c.JSON(500, gin.H{
			"code":    500,
			"message": "Internal Server Error",
			"error":   fmt.Sprintf("%v", err),
		})
		c.Abort()
	}
}
