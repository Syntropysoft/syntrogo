package security

import (
	"strings"
	"github.com/syntropysoft/syntrogo/src/domain"
)

// BearerToken middleware validates Bearer tokens.
// Usage: app.Use(BearerToken("secret123"))
func BearerToken(expectedToken string) domain.Middleware {
	return func(next domain.HandlerFunc) domain.HandlerFunc {
		return func(ctx *domain.Context) error {
			// Extract Authorization header
			auth := ctx.Header("Authorization")
			if auth == "" {
				return domain.NewHTTPException(401, "Missing Authorization header")
			}

			// Check Bearer prefix
			if !strings.HasPrefix(auth, "Bearer ") {
				return domain.NewHTTPException(401, "Invalid Authorization header format")
			}

			// Extract token
			token := strings.TrimPrefix(auth, "Bearer ")
			if token == "" {
				return domain.NewHTTPException(401, "Missing token")
			}

			// Validate token
			if token != expectedToken {
				return domain.NewHTTPException(401, "Invalid token")
			}

			// Set user in context (can be extended)
			ctx.Headers["X-Authenticated-User"] = "user"

			// Continue to next handler
			return next(ctx)
		}
	}
}

