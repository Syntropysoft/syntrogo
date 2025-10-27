package infrastructure

// Infrastructure package contains adapters to external systems.
//
// Adapters:
// - HTTPAdapter: Adapts net/http to our domain
// - Future: Redis, Database, etc.
//
// Principles:
// - Adapter Pattern: Wraps external libraries
// - Dependency Inversion: Depends on abstractions (domain)
// - Isolation: Infrastructure changes don't affect domain

