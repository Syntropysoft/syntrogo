package infrastructure

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/syntropysoft/syntrogo/src/application"
	"github.com/syntropysoft/syntrogo/src/domain"
)

// HTTPContext extends domain.Context with HTTP-specific data.
// Composition pattern: extends base Context
type HTTPContext struct {
	*domain.Context
	Request  *http.Request
	Response http.ResponseWriter
}

// HTTPAdapter adapts net/http to our domain.
// Implements http.Handler from net/http
type HTTPAdapter struct {
	routeRegistry      *application.RouteRegistry
	middlewareRegistry *application.MiddlewareRegistry
	validator          *validator.Validate
	swaggerEnabled     bool
	swaggerSpec        map[string]interface{}
}

// NewHTTPAdapter creates a new HTTP adapter.
func NewHTTPAdapter(routeRegistry *application.RouteRegistry, middlewareRegistry *application.MiddlewareRegistry) *HTTPAdapter {
	return &HTTPAdapter{
		routeRegistry:      routeRegistry,
		middlewareRegistry: middlewareRegistry,
		validator:          validator.New(),
	}
}

// ServeHTTP implements http.Handler interface.
// This handles all HTTP requests.
func (a *HTTPAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle Swagger endpoint
	if r.URL.Path == "/swagger.json" && a.swaggerEnabled {
		a.handleSwagger(w)
		return
	}

	// Find route
	route := a.routeRegistry.Find(r.Method, r.URL.Path)
	
	// If not found
	if route == nil {
		http.NotFound(w, r)
		return
	}
	
	// Create domain context
	ctx := &domain.Context{
		Params:      make(map[string]string),
		QueryParams: make(map[string]string),
		Headers:     make(map[string]string),
		StatusCode:  200,
	}
	
	// Bind query parameters
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			ctx.QueryParams[key] = values[0]
		}
	}
	
	// Bind headers
	for key, values := range r.Header {
		if len(values) > 0 {
			ctx.Headers[key] = values[0]
		}
	}
	
	// Store request and response in context for BindJSON
	ctx.Request = r
	ctx.Response = w
	ctx.Binder = a // Set binder for BindJSON
	
	// Apply global middlewares
	handler := route.Handler
	handler = a.middlewareRegistry.Apply(handler)
	
	// Apply route-specific middlewares
	for i := len(route.Middlewares) - 1; i >= 0; i-- {
		handler = route.Middlewares[i](handler)
	}
	
	// Call handler with all middlewares applied
	if err := handler(ctx); err != nil {
		a.handleError(w, err)
		return
	}
	
	// Write response
	if ctx.Body != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(ctx.StatusCode)
		json.NewEncoder(w).Encode(ctx.Body)
	} else if ctx.StatusCode != 200 {
		w.WriteHeader(ctx.StatusCode)
	}
}

// handleError handles errors from handlers.
func (a *HTTPAdapter) handleError(w http.ResponseWriter, err error) {
	// If it's our HTTPException, use its status code
	if httpErr, ok := err.(*domain.HTTPException); ok {
		w.WriteHeader(httpErr.StatusCode)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": httpErr.Message,
		})
		return
	}
	
	// Generic error
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": "Internal Server Error",
	})
}

// StartServer starts the HTTP server on the given port.
func (a *HTTPAdapter) StartServer(port string) error {
	return http.ListenAndServe(":"+port, a)
}

// SetSwaggerEnabled enables Swagger documentation.
func (a *HTTPAdapter) SetSwaggerEnabled(enabled bool) {
	a.swaggerEnabled = enabled
}

// SetSwaggerSpec sets the OpenAPI spec to serve.
func (a *HTTPAdapter) SetSwaggerSpec(spec map[string]interface{}) {
	a.swaggerSpec = spec
}

// handleSwagger serves the OpenAPI JSON specification.
func (a *HTTPAdapter) handleSwagger(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	if a.swaggerSpec == nil {
		http.Error(w, "Swagger spec not available", 500)
		return
	}

	json.NewEncoder(w).Encode(a.swaggerSpec)
}

// BindJSON decodes JSON from the request body and validates it.
func (a *HTTPAdapter) BindJSON(ctx *domain.Context, v interface{}) error {
	if ctx.Request == nil {
		return domain.NewHTTPException(400, "request not available")
	}
	
	// Type assert to *http.Request
	req, ok := ctx.Request.(*http.Request)
	if !ok {
		return domain.NewHTTPException(500, "invalid request type")
	}
	
	// Read body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return domain.NewHTTPException(400, "failed to read request body")
	}
	
	// Decode JSON
	if len(body) > 0 {
		if err := json.Unmarshal(body, v); err != nil {
			return domain.NewHTTPException(400, "invalid JSON")
		}
	}
	
	// Validate
	if err := a.validator.Struct(v); err != nil {
		return domain.NewHTTPException(422, err.Error())
	}
	
	return nil
}

