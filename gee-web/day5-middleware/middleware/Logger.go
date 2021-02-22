package middleware

import (
	"7days-golang/gee-web/day5-middleware/gee"
	"log"
	"time"
)

func Logger() gee.HandlerFunc  {
	return func(c *gee.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(start))
	}
}
