package middlewares

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	tokens       chan struct{}
	refillTicker *time.Ticker
	stopChan     chan struct{}
}

//IP Rate Limiter specific code <Starts>

type IPRateLimiter struct {
	mu         sync.Mutex
	limiters   map[string]*RateLimiter
	capacity   int
	refillTime time.Duration
}

func NewIPRateLimiter(capacity int, refillInterval time.Duration) *IPRateLimiter {
	return &IPRateLimiter{
		limiters:   make(map[string]*RateLimiter),
		capacity:   capacity,
		refillTime: refillInterval,
	}
}

func (ipr *IPRateLimiter) getLimiter(ip string) *RateLimiter {
	ipr.mu.Lock()
	defer ipr.mu.Unlock()

	if limiter, exists := ipr.limiters[ip]; exists {
		return limiter
	}

	limiter := NewRateLimiter(ipr.capacity, ipr.refillTime)
	ipr.limiters[ip] = limiter
	return limiter
}

// getClientIP extracts client IP from context
func getClientIP(c *gin.Context) string {
	ip := c.ClientIP()
	parsedIP := net.ParseIP(ip)
	if parsedIP != nil {
		return parsedIP.String()
	}
	return "unknown"
}

//IP Rate Limiter specific code <Ends>

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

func (ipr *IPRateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIP(c)
		limiter := ipr.getLimiter(ip)

		select {
		case <-limiter.tokens:
			c.Next()
		default:
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
		}
	}
}

func (rl *RateLimiter) Stop() {
	close(rl.stopChan)
}
