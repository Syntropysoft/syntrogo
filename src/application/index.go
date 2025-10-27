// Package application contains use cases and application logic.
// DDD: Application layer - orchestrates domain logic without external dependencies
package application

// Modules:
// - RouteRegistry: Manages HTTP routes
// - SchemaValidator: Validates structs with go-playground/validator
// - OpenAPIGenerator: Generates OpenAPI 3.0 specs from reflection
// - MiddlewareRegistry: Manages middleware chain
//
// Principles:
// - SOLID: Each module has single responsibility
// - DDD: Orchestrates domain, depends on abstractions
// - Adapter Pattern: Wraps external libraries (validator, etc.)
// - Guard Clauses: Validates early, fails fast
// - Functional: Composes functions (middleware chain)

