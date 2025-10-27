package application

import (
	"github.com/syntropysoft/syntrogo/src/domain"
)

// RouteRegistry manages HTTP routes.
// SOLID: Single Responsibility - only registers routes
// DDD: Use case in Application layer
type RouteRegistry struct {
	routes []*domain.Route
}

// NewRouteRegistry creates a new route registry.
func NewRouteRegistry() *RouteRegistry {
	return &RouteRegistry{
		routes: make([]*domain.Route, 0),
	}
}

// Register adds a new route to the registry.
// Guard Clause: Validates route before adding (SOLID: Single Responsibility)
func (r *RouteRegistry) Register(method, path string, handler domain.HandlerFunc, options domain.RouteOptions) error {
	// Guard clause: Fail fast validation
	if path == "" {
		return domain.NewHTTPException(400, "path is required")
	}
	if handler == nil {
		return domain.NewHTTPException(400, "handler is required")
	}

	// Happy path: Create and register route
	route := &domain.Route{
		Method:      method,
		Path:        path,
		Handler:     handler,
		Options:     options,
		Middlewares: options.Middlewares,
	}
	
	r.routes = append(r.routes, route)
	return nil
}

// Find returns a route matching the method and path.
// Returns nil if not found.
func (r *RouteRegistry) Find(method, path string) *domain.Route {
	for _, route := range r.routes {
		if route.Method == method {
			// TODO: Add pattern matching for path parameters like /users/:id
			// For now, exact match only
			if route.Path == path {
				return route
			}
		}
	}
	return nil
}

// GetRoutes returns all registered routes.
func (r *RouteRegistry) GetRoutes() []*domain.Route {
	return r.routes
}

