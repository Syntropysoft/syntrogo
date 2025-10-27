package domain

// Route represents an HTTP route with its handler and options.
// This is a pure domain entity with no external dependencies.
type Route struct {
	Method     string
	Path       string
	Handler    HandlerFunc
	Options    RouteOptions
	Middlewares []Middleware  // Middlewares specific to this route
}

// HandlerFunc is the function signature for route handlers.
// Controllers are functions that receive a Context and return an error.
type HandlerFunc func(*Context) error

// Context provides access to the current HTTP request and response.
// Pure value object, no external dependencies.
type Context struct {
	Request      interface{} // *http.Request (set by infrastructure)
	Response     interface{} // http.ResponseWriter (set by infrastructure)
	Params       map[string]string
	QueryParams  map[string]string
	Headers      map[string]string
	StatusCode   int
	Body         interface{}
	
	// Infrastructure adapter for BindJSON (set by infrastructure)
	Binder       Binder
}

// Binder interface allows Context to bind JSON without knowing the implementation.
type Binder interface {
	BindJSON(*Context, interface{}) error
}

// RouteOptions contains additional metadata for a route.
type RouteOptions struct {
	Body       interface{}          // Request body type
	Response   interface{}          // Response type
	Summary    string               // Endpoint summary
	Tags       []string             // OpenAPI tags
	Params     map[string]ParamSpec  // Path parameters
	Middlewares []Middleware        // Middlewares for this route
}

// ParamSpec specifies validation rules for path parameters.
type ParamSpec struct {
	Type    string // "string", "integer", etc.
	Required bool
	Min      *int
	Max      *int
}

// Param returns a path parameter by name.
func (c *Context) Param(name string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[name]
}

// Query returns a query parameter by name.
func (c *Context) Query(name string) string {
	if c.QueryParams == nil {
		return ""
	}
	return c.QueryParams[name]
}

// Header returns a header value by name.
func (c *Context) Header(name string) string {
	if c.Headers == nil {
		return ""
	}
	return c.Headers[name]
}

// SetHeader sets a response header.
func (c *Context) SetHeader(name, value string) {
	if c.Headers == nil {
		c.Headers = make(map[string]string)
	}
	c.Headers[name] = value
}

// Status sets the response status code.
func (c *Context) Status(code int) *Context {
	c.StatusCode = code
	return c
}

// BindJSON binds the request body to a struct and validates it.
func (c *Context) BindJSON(v interface{}) error {
	if c.Binder == nil {
		return NewHTTPException(500, "binder not available")
	}
	return c.Binder.BindJSON(c, v)
}

// JSON writes a JSON response.
// This will be implemented by the infrastructure layer.
func (c *Context) JSON(statusCode int, data interface{}) error {
	c.StatusCode = statusCode
	c.Body = data
	// This will be implemented by infrastructure
	return nil
}

