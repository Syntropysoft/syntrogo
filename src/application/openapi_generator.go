package application

import (
	"reflect"
	"strings"

	"github.com/syntropysoft/syntrogo/src/domain"
)

// OpenAPIGenerator generates OpenAPI 3.0 specifications from routes.
// SOLID: Single Responsibility - only generates OpenAPI
// Reflection-based: Reads struct tags to infer schemas
type OpenAPIGenerator struct {
	routes []*domain.Route
}

// NewOpenAPIGenerator creates a new OpenAPI generator.
func NewOpenAPIGenerator(routes []*domain.Route) *OpenAPIGenerator {
	return &OpenAPIGenerator{
		routes: routes,
	}
}

// Generate creates the OpenAPI 3.0 specification.
// Uses reflection to infer schemas from struct tags
func (g *OpenAPIGenerator) Generate(title, version string) (map[string]interface{}, error) {
	spec := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":   title,
			"version": version,
		},
		"paths": map[string]interface{}{},
	}

	// Add each route to the spec
	for _, route := range g.routes {
		pathSpec := spec["paths"].(map[string]interface{})
		pathKey := strings.ToLower(route.Method) + " " + route.Path
		
		pathSpec[pathKey] = g.generatePathItem(route)
	}

	return spec, nil
}

// generatePathItem creates the OpenAPI path item for a route.
func (g *OpenAPIGenerator) generatePathItem(route *domain.Route) map[string]interface{} {
	item := map[string]interface{}{
		"summary": route.Options.Summary,
		"tags":    route.Options.Tags,
	}

	// Add request body if defined
	if route.Options.Body != nil {
		item["requestBody"] = map[string]interface{}{
			"required": true,
			"content": map[string]interface{}{
				"application/json": map[string]interface{}{
					"schema": g.inferSchemaFromStruct(route.Options.Body),
				},
			},
		}
	}

	// Add response if defined
	if route.Options.Response != nil {
		item["responses"] = map[string]interface{}{
			"200": map[string]interface{}{
				"description": "Success",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": g.inferSchemaFromStruct(route.Options.Response),
					},
				},
			},
		}
	}

	return item
}

// addRouteToSpec adds a route to the OpenAPI specification.
func (g *OpenAPIGenerator) addRouteToSpec(spec map[string]interface{}, route *domain.Route) {
	// TODO: Add route to OpenAPI spec
	// Uses reflection to infer body/response schemas
}

// inferSchemaFromStruct uses reflection to infer OpenAPI schema from struct tags.
func (g *OpenAPIGenerator) inferSchemaFromStruct(structType interface{}) map[string]interface{} {
	t := reflect.TypeOf(structType)
	
	// If it's a pointer, get the element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		// For non-structs, return basic type
		return map[string]interface{}{
			"type": g.goTypeToOpenAPIType(t.Kind().String()),
		}
	}

	// For slices/arrays
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return map[string]interface{}{
			"type": "array",
			"items": g.inferSchemaFromStruct(t.Elem()),
		}
	}

	properties := map[string]interface{}{}
	required := []string{}

	// Iterate through struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		
		// Get json tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Remove omitempty
		jsonTag = strings.Split(jsonTag, ",")[0]

		// Infer property type
		prop := map[string]interface{}{
			"type": g.goTypeToOpenAPIType(field.Type.Kind().String()),
		}

		// Check for validation tags
		validateTag := field.Tag.Get("validate")
		if strings.Contains(validateTag, "required") {
			required = append(required, jsonTag)
		}

		properties[jsonTag] = prop
	}

	schema := map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}

	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

// goTypeToOpenAPIType converts GO types to OpenAPI types.
func (g *OpenAPIGenerator) goTypeToOpenAPIType(goType string) string {
	switch goType {
	case "string":
		return "string"
	case "int", "int8", "int16", "int32", "int64":
		return "integer"
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	default:
		return "string"
	}
}

