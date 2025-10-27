package api

import "github.com/syntropysoft/syntrogo/src/domain"

// Body specifies the request body type.
// Use as: Body(UserRequest{})
func Body(typ interface{}) domain.RouteOptions {
	// TODO: Return route option for body
	return domain.RouteOptions{}
}

// Response specifies the response type and status code.
// Use as: Response(201, UserResponse{})
func Response(statusCode int, typ interface{}) domain.RouteOptions {
	// TODO: Return route option for response
	return domain.RouteOptions{}
}

// Summary sets the endpoint summary for Swagger.
// Use as: Summary("Create user")
func Summary(text string) domain.RouteOptions {
	// TODO: Return route option for summary
	return domain.RouteOptions{}
}

// Tags sets the OpenAPI tags for Swagger.
// Use as: Tags("users", "admin")
func Tags(tags ...string) domain.RouteOptions {
	// TODO: Return route option for tags
	return domain.RouteOptions{}
}

// Params specifies path parameter validation.
// Use as: Params(map[string]string{"id": "integer"})
func Params(spec map[string]string) domain.RouteOptions {
	// TODO: Return route option for params
	return domain.RouteOptions{}
}
