package main

import "time"

package main

import (
"time"

"github.com/gin-gonic/gin"

"yourmodule/middleware" // replace with your actual module path
)

func main() {
	r := gin.Default()

	// ==== Option 1: Global Rate Limiter ====
	globalLimiter := middleware.NewRateLimiter(10, time.Second)
	r.Use(globalLimiter.Middleware())

	// ==== Option 2: Per-IP Rate Limiter ====
	// perIPLimiter := middleware.NewIPRateLimiter(5, 2*time.Second)
	// r.Use(perIPLimiter.Middleware())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")

	defer globalLimiter.Stop()

	// or
	// defer perIPLimiter.Stop() for each limiter in the map
}
