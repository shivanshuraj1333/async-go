package main

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	mu          sync.Mutex
	window      time.Duration
	maxReq      int
	requests    int
	windowStart time.Time
}

func NewFixedWindowLimiter(window time.Duration, maxReq int) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		window:      window,
		maxReq:      maxReq,
		windowStart: time.Now(),
	}
}

func (r *FixedWindowLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.Sub(r.windowStart) >= r.window {
		r.windowStart = now
		r.requests = 0
	}
	if r.requests < r.maxReq {
		r.requests++
		return true
	}
	return false
}
