package application

import (
	"github.com/syntropysoft/syntrogo/src/domain"
)

// MiddlewareRegistry manages middleware functions.
// SOLID: Single Responsibility - only manages middleware
type MiddlewareRegistry struct {
	middlewares []domain.Middleware
}

// NewMiddlewareRegistry creates a new middleware registry.
func NewMiddlewareRegistry() *MiddlewareRegistry {
	return &MiddlewareRegistry{
		middlewares: make([]domain.Middleware, 0),
	}
}

// Use adds a middleware to the chain.
func (r *MiddlewareRegistry) Use(mw domain.Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

// Apply applies all middleware to a handler in order.
// Functional: Composes middleware functions
func (r *MiddlewareRegistry) Apply(handler domain.HandlerFunc) domain.HandlerFunc {
	// Functional composition: Chain middleware
	result := handler
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		result = r.middlewares[i](result)
	}
	return result
}

