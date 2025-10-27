package security

import (
	"github.com/syntropysoft/syntrogo/src/domain"
)

// APIKey middleware validates API keys.
// Usage: app.Use(APIKey("sk_live_..."))
func APIKey(expectedKey string) domain.Middleware {
	return func(next domain.HandlerFunc) domain.HandlerFunc {
		return func(ctx *domain.Context) error {
			// Extract API key from header
			apiKey := ctx.Header("X-API-Key")
			if apiKey == "" {
				return domain.NewHTTPException(401, "Missing X-API-Key header")
			}

			// Validate key
			if apiKey != expectedKey {
				return domain.NewHTTPException(401, "Invalid API key")
			}

			// Continue to next handler
			return next(ctx)
		}
	}
}

