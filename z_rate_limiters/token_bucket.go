package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	tokens       chan struct{}
	refillTicker *time.Ticker
	stopChan     chan struct{}
}

// NewRateLimiter creates a token bucket with capacity and refill interval.
func NewRateLimiter(capacity int, refillInterval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:       make(chan struct{}, capacity),
		refillTicker: time.NewTicker(refillInterval),
		stopChan:     make(chan struct{}),
	}

	// Fill initial tokens
	for i := 0; i < capacity; i++ {
		rl.tokens <- struct{}{}
	}

	// Start refilling tokens
	go rl.refill()

	return rl
}

func (rl *RateLimiter) refill() {
	for {
		select {
		case <-rl.refillTicker.C:
			select {
			case rl.tokens <- struct{}{}:
			default: // bucket full, skip
			}
		case <-rl.stopChan:
			rl.refillTicker.Stop()
			return
		}
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		select {
		case <-rl.tokens:
			c.Next()
		default:
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
		}
	}
}

func (rl *RateLimiter) Stop() {
	close(rl.stopChan)
}
