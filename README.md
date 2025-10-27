# SyntroGo

<div align="center">
  <img src="assets/syntropySoft.png" alt="SyntroGo" width="400" />
</div>

<div align="center">

**FastAPI para GO: Simplicidad de Tesla con optimización de Ferrari**

[![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square)](LICENSE)
[![GitHub](https://img.shields.io/badge/github-Syntropysoft/syntrogo-181717?style=flat-square&logo=github)](https://github.com/Syntropysoft/syntrogo)

> Part of the [SyntropySoft](https://syntropysoft.com) ecosystem

</div>

## 🎯 ¿Qué es SyntroGo?

**FastAPI para GO: Simplicidad de Tesla con optimización de Ferrari.**

La librería GO más simple para construir APIs, con todo el poder nativo de GO (goroutines, compilación, reflection).

### Visión

> **"Usa el framework como si fuera un Tesla, pero obtén el rendimiento de un Ferrari"**

**Simplicidad:** API fluente como FastAPI (tesla-level simplicity)  
**Performance:** Optimización nativa de GO (ferrari/mclaren-level speed)

## 🚀 Quick Start

```go
package main

import api "github.com/syntropysoft/syntrogo"

type UserRequest struct {
    Name string `json:"name" validate:"required,min=3"`
}

type UserResponse struct {
    Message string `json:"message"`
}

func handler(c *api.Context) error {
    var req UserRequest
    if err := c.BindJSON(&req); err != nil {
        return err
    }
    return c.JSON(200, UserResponse{Message: "Hello " + req.Name})
}

func main() {
    app := api.New().Title("My API").REST()
    app.POST("/hello", handler, api.Body(UserRequest{}))
    app.Listen(":3000")
}
```

## 🎯 Features

- ✅ **API Fluent** - Sintaxis declarativa como FastAPI
- ✅ **Swagger Automático** - Zero manual comments
- ✅ **Validation Integrada** - Go struct tags
- ✅ **Testing Simple** - TinyTest para tests rápidos
- ✅ **Performance Nativo** - Goroutines & compilación GO
- ✅ **Protocol Agnostic** - REST (v1.0) + gRPC (v2.0+)

## 🏗️ Estructura del Proyecto

```txt
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

## 🧠 Principios de Diseño

- **SOLID** - Single Responsibility en cada módulo
- **DDD** - Separación Domain/Application/Infrastructure
- **Guard Clauses** - Fail fast validation
- **Functional** - Composición de funciones
- **Concurrency First** - Goroutines nativas de GO

## 📖 Documentación

- [Filosofía](./PHILOSOPHY_GO.md) - Principios y visión
- [Arquitectura](../ARCHITECTURE_GO.md) - Diseño y patrones
- [Testing](./TESTING_STRATEGY.md) - TinyTest & SmartMutator
- [Ejemplos](../syntrogo-examples/) - Ejemplos de uso

## 🔗 Links

- [Repositorio GitHub](https://github.com/Syntropysoft/syntrogo) - Código fuente
- [SyntropySoft](https://syntropysoft.com) - Ecosystem
- [Documentación](./docs/) - Docs completas (coming soon)

## 📄 License

Apache 2.0 - Part of the SyntropySoft ecosystem
