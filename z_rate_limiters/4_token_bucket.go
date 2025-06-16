package main

import (
	"sync"
	"time"
)

type SlidingWindowLog struct {
	mu       sync.Mutex
	window   time.Duration
	maxReq   int
	requests []time.Time
}

func NewSlidingWindowLog(window time.Duration, maxReq int) *SlidingWindowLog {
	return &SlidingWindowLog{
		window:   window,
		maxReq:   maxReq,
		requests: []time.Time{},
	}
}

func (r *SlidingWindowLog) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-r.window)

	// Remove old requests
	var filtered []time.Time
	for _, t := range r.requests {
		if t.After(cutoff) {
			filtered = append(filtered, t)
		}
	}
	r.requests = filtered

	if len(r.requests) < r.maxReq {
		r.requests = append(r.requests, now)
		return true
	}
	return false
}
