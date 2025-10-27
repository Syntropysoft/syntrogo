// Package security provides security middlewares.
//
// Middlewares available:
// - BearerToken: Validates Bearer token in Authorization header
// - APIKey: Validates API key in X-API-Key header
// - CORS: Handles CORS headers
// - RateLimit: Applies rate limiting
//
// Usage:
//   import "github.com/syntropysoft/syntrogo/src/security"
//
//   app.Use(security.BearerToken("token123"))
//   app.Use(security.APIKey("key123"))
//   app.Use(security.CORS("*"))
//   app.Use(security.RateLimit(limiter))
package security

