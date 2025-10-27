// Package domain contains pure domain entities with no external dependencies.
// These entities follow DDD principles: no dependencies on infrastructure or application layers.
package domain

// Exports:
// - Route: HTTP route entity
// - Context: Request context value object
// - HTTPException: Domain exceptions
// - Types: Middleware, AppConfig, etc.
//
// Principles:
// - SOLID: Single Responsibility (each entity has one purpose)
// - DDD: Pure domain, zero external dependencies
// - Guard Clauses: Fail fast validation
// - Functional: Immutable where possible

