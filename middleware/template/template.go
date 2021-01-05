package template

import (
	"github.com/gin-gonic/gin"
)

func MiddlewareTemplate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// DO Something

		c.Next()
	}
}
