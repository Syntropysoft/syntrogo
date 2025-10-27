package security

import (
	"github.com/syntropysoft/syntrogo/src/domain"
)

// CORS middleware handles CORS headers.
// Usage: app.Use(CORS("*")) or app.Use(CORS("https://example.com"))
func CORS(allowedOrigin string) domain.Middleware {
	return func(next domain.HandlerFunc) domain.HandlerFunc {
		return func(ctx *domain.Context) error {
			// Set CORS headers
			ctx.SetHeader("Access-Control-Allow-Origin", allowedOrigin)
			ctx.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			ctx.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key")
			ctx.SetHeader("Access-Control-Allow-Credentials", "true")

			// Check if this is OPTIONS request (preflight)
			method := ctx.Header(":method")
			if method == "OPTIONS" {
				ctx.Status(204)
				return nil
			}

			// Continue to next handler
			return next(ctx)
		}
	}
}

