# SyntroGo

> **FastAPI para GO: Simplicidad de Tesla con optimización de Ferrari**

> Part of the [SyntropySoft](https://syntropysoft.com) ecosystem

## 🎯 Visión

Crear la librería más simple y potente para construir APIs en GO:
- **Simplicidad:** API fluente como FastAPI (tesla-level simplicity)
- **Performance:** Optimización nativa de GO (ferrari/mclaren-level speed)

## 🏗️ Estructura del Proyecto

```
syntrogo/
├── src/
│   ├── domain/          # Entidades puras (no dependencias externas)
│   │   ├── route.go
│   │   ├── types.go
│   │   └── httpexception.go
│   │
│   ├── application/     # Use cases (orquesta lógica)
│   │   ├── route_registry.go
│   │   ├── schema_validator.go
│   │   ├── openapi_generator.go
│   │   └── middleware_registry.go
│   │
│   ├── core/            # Framework principal (API pública)
│   │   └── app.go
│   │
│   ├── infrastructure/  # Adapters (net/http, etc.)
│   │   └── http_adapter.go
│   │
│   └── testing/         # Testing utilities
│       ├── tiny_test.go
│       └── smart_mutator.go
│
└── syntrogo.go        # API pública (entry point)
```

## 🎯 Principios

- **SOLID** - Single Responsibility en cada módulo
- **DDD** - Separación Domain/Application/Infrastructure
- **Guard Clauses** - Fail fast validation
- **Functional** - Composición de funciones
- **Concurrency First** - Goroutines nativas de GO

## 🚀 Estado

**Maqueta inicial creada** - Declaraciones de intenciones en lugar de implementación.

Próximos pasos:
1. Implementar Core Fluent API
2. Implementar Validation con go-playground/validator
3. Implementar Swagger generation con reflection
4. Implementar TinyTest
5. Implementar SmartMutator

