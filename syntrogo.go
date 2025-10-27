// Package syntrogo provides the main public API.
//
// This is the entry point for users: import "github.com/syntropysoft/syntrogo"
//
// # Part of the SyntropySoft ecosystem
//
// Philosophy:
// - Tesla-Level Simplicity: Simple, intuitive API
// - Ferrari-Level Performance: Optimized under the hood
//
// Example usage:
//
//	import api "github.com/syntropysoft/syntrogo"
//
//	app := api.New()
//	app.POST("/users", handler,
//	    api.Body(UserRequest{}),
//	    api.Response(201, UserResponse{}),
//	)
//	app.Listen(":3000")
//
// Using alias 'api' makes the code more fluent and readable.
package syntrogo

import (
	"github.com/syntropysoft/syntrogo/src/core"
	"github.com/syntropysoft/syntrogo/src/domain"
)

// New creates a new application instance.
// This is the main entry point for users
func New() *core.App {
	return core.New()
}

// Re-export domain types for user convenience
type (
	Context     = domain.Context
	Handler     = domain.HandlerFunc
	RouteOptions = domain.RouteOptions
)

// Context represents the request context.
// Re-exported from domain for user convenience

// Helper functions for route options

// Body specifies the request body type.
func Body(typ interface{}) RouteOptions {
	return RouteOptions{Body: typ}
}

// Response specifies the response type and status code.
func Response(statusCode int, typ interface{}) RouteOptions {
	return RouteOptions{Response: typ}
}

// Summary sets the endpoint summary for Swagger.
func Summary(text string) RouteOptions {
	return RouteOptions{Summary: text}
}

// Tags sets the OpenAPI tags for Swagger.
func Tags(tags ...string) RouteOptions {
	return RouteOptions{Tags: tags}
}

// Params specifies path parameter validation.
func Params(spec map[string]string) RouteOptions {
	// Convert map to ParamSpec map
	paramSpec := make(map[string]domain.ParamSpec)
	for key, value := range spec {
		// TODO: Parse the value string to get type, required, etc.
		paramSpec[key] = domain.ParamSpec{
			Type:     value,
			Required: true,
		}
	}
	return RouteOptions{Params: paramSpec}
}

// Middleware applies a middleware to this specific route.
func Middleware(mw domain.Middleware) RouteOptions {
	return RouteOptions{Middlewares: []domain.Middleware{mw}}
}

