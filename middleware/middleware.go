package middleware

import "github.com/gin-gonic/gin"

type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *GoMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Next()
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
