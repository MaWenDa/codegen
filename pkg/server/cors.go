package server

import "github.com/gin-gonic/gin"

// Cors 解决跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = "*"
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
