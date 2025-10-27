package domain

// Middleware is a function that wraps a Handler.
// Used for cross-cutting concerns like logging, authentication, etc.
type Middleware func(HandlerFunc) HandlerFunc

// AppConfig holds the application configuration.
type AppConfig struct {
	Title       string
	Version     string
	Swagger     bool
	SwaggerPath string
	Port        string
}

