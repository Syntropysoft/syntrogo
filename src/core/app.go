package core

import (
	"github.com/syntropysoft/syntrogo/src/application"
	"github.com/syntropysoft/syntrogo/src/domain"
	"github.com/syntropysoft/syntrogo/src/infrastructure"
)

// These are placeholders - will be implemented
type RouteRegistry = application.RouteRegistry
type MiddlewareRegistry = application.MiddlewareRegistry

// Protocol defines the communication protocol.
type Protocol int

const (
	// ProtocolREST represents REST/HTTP protocol.
	ProtocolREST Protocol = iota
	// ProtocolGRPC represents gRPC protocol (future).
	ProtocolGRPC
	// ProtocolBoth runs both REST and gRPC (future).
	ProtocolBoth
)

// App is the main application instance.
// This is the public API of the framework (like SyntroJS.App or FastAPI.Application)
// It's protocol-agnostic: can run REST, gRPC, or both.
type App struct {
	config             *domain.AppConfig
	routeRegistry      *RouteRegistry
	middlewareRegistry *MiddlewareRegistry
	swaggerEnabled     bool
	protocol           Protocol           // Protocol selector (REST by default, gRPC for v2.0+)
	prefix             string            // Route prefix for groups
	groupMiddlewares   []domain.Middleware // Middlewares for this group
	parent             *App              // Parent app for groups
}

// New creates a new application instance.
// This is the entry point for users: simplicity.New()
func New() *App {
	return &App{
		config: &domain.AppConfig{
			Title:   "SyntroGo API",
			Version: "1.0.0",
			Port:    "3000",
		},
		routeRegistry:     application.NewRouteRegistry(),
		middlewareRegistry: application.NewMiddlewareRegistry(),
		swaggerEnabled:    false,
		protocol:          ProtocolREST, // Default to REST
	}
}

// REST sets the protocol to REST (HTTP).
// This is the default protocol for v1.0.
func (a *App) REST() *App {
	a.protocol = ProtocolREST
	return a
}

// GRPC sets the protocol to gRPC.
// Available in v2.0+
func (a *App) GRPC() *App {
	a.protocol = ProtocolGRPC
	return a
}

// GET registers a GET route.
// API Fluent: Returns *App for chaining
// Can chain: app.GET(...).POST(...)
func (a *App) GET(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *App {
	a.registerRoute("GET", path, handler, opts...)
	return a
}

// POST registers a POST route.
func (a *App) POST(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *App {
	a.registerRoute("POST", path, handler, opts...)
	return a
}

// PUT registers a PUT route.
func (a *App) PUT(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *App {
	a.registerRoute("PUT", path, handler, opts...)
	return a
}

// DELETE registers a DELETE route.
func (a *App) DELETE(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *App {
	a.registerRoute("DELETE", path, handler, opts...)
	return a
}

// registerRoute is the internal implementation that merges options.
func (a *App) registerRoute(method, path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) {
	// Merge all options into one
	var merged domain.RouteOptions
	for _, opt := range opts {
		if opt.Body != nil {
			merged.Body = opt.Body
		}
		if opt.Response != nil {
			merged.Response = opt.Response
		}
		if opt.Summary != "" {
			merged.Summary = opt.Summary
		}
		if len(opt.Tags) > 0 {
			merged.Tags = opt.Tags
		}
		if opt.Params != nil {
			merged.Params = opt.Params
		}
		if len(opt.Middlewares) > 0 {
			merged.Middlewares = opt.Middlewares
		}
	}
	
	// Combine group middlewares with route middlewares
	merged.Middlewares = append(a.groupMiddlewares, merged.Middlewares...)
	
	// Register in registry with full path (prefix + path)
	fullPath := a.prefix + path
	_ = a.routeRegistry.Register(method, fullPath, handler, merged)
}

// Use adds a global middleware to the chain.
func (a *App) Use(middleware domain.Middleware) *App {
	a.middlewareRegistry.Use(middleware)
	return a
}

// Group creates a route group with a prefix and optional middlewares.
// Usage: api := app.Group("/api", security.BearerToken("token"))
func (a *App) Group(prefix string, middlewares ...domain.Middleware) *App {
	return &App{
		config:           a.config,
		routeRegistry:    a.routeRegistry,
		middlewareRegistry: a.middlewareRegistry,
		swaggerEnabled:   a.swaggerEnabled,
		protocol:         a.protocol,
		prefix:           prefix,
		groupMiddlewares: middlewares,
		parent:           a,
	}
}

// Title sets the application title.
func (a *App) Title(title string) *App {
	a.config.Title = title
	return a
}

// Version sets the application version.
func (a *App) Version(version string) *App {
	a.config.Version = version
	return a
}

// Swagger enables Swagger documentation.
func (a *App) Swagger(enabled bool) *App {
	a.swaggerEnabled = enabled
	return a
}

// Listen starts the HTTP server.
// The app is now ready to accept requests
func (a *App) Listen(port string) error {
	a.config.Port = port
	
	// Create HTTP adapter
	adapter := infrastructure.NewHTTPAdapter(
		a.routeRegistry,
		a.middlewareRegistry,
	)
	
	// Generate and set Swagger if enabled
	if a.swaggerEnabled {
		generator := application.NewOpenAPIGenerator(a.routeRegistry.GetRoutes())
		spec, err := generator.Generate(a.config.Title, a.config.Version)
		if err == nil {
			adapter.SetSwaggerEnabled(true)
			adapter.SetSwaggerSpec(spec)
		}
	}
	
	// Start server
	return adapter.StartServer(port)
}

// GetRouteRegistry returns the route registry (for testing).
func (a *App) GetRouteRegistry() *RouteRegistry {
	return a.routeRegistry
}

// GetMiddlewareRegistry returns the middleware registry (for testing).
func (a *App) GetMiddlewareRegistry() *MiddlewareRegistry {
	return a.middlewareRegistry
}

