package security

import (
	"sync"
	"time"
	"github.com/syntropysoft/syntrogo/src/domain"
)

// RateLimiter is a simple in-memory rate limiter.
type RateLimiter struct {
	requests map[string][]time.Time
	max      int
	duration time.Duration
	mu       sync.Mutex
}

// NewRateLimiter creates a new rate limiter.
// max: maximum number of requests
// duration: time window (e.g., 1 minute)
func NewRateLimiter(max int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		max:      max,
		duration: duration,
	}
}

// RateLimit middleware applies rate limiting.
// Usage: app.Use(RateLimit(limiter))
func RateLimit(limiter *RateLimiter) domain.Middleware {
	return func(next domain.HandlerFunc) domain.HandlerFunc {
		return func(ctx *domain.Context) error {
			// Get client IP
			clientIP := ctx.Header("X-Forwarded-For")
			if clientIP == "" {
				clientIP = ctx.Header("X-Real-IP")
			}

			// Default IP if not found
			if clientIP == "" {
				clientIP = "unknown"
			}

			// Check rate limit
			if !limiter.Allow(clientIP) {
				return domain.NewHTTPException(429, "Too many requests")
			}

			// Continue to next handler
			return next(ctx)
		}
	}
}

// Allow checks if a request is allowed.
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	
	// Clean old requests
	if times, ok := rl.requests[key]; ok {
		valid := []time.Time{}
		for _, t := range times {
			if now.Sub(t) < rl.duration {
				valid = append(valid, t)
			}
		}
		rl.requests[key] = valid
		
		// Check if limit exceeded
		if len(valid) >= rl.max {
			return false
		}
	}

	// Add new request
	rl.requests[key] = append(rl.requests[key], now)
	return true
}

