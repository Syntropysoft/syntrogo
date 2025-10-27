package testing

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/syntropysoft/syntrogo/src/application"
	"github.com/syntropysoft/syntrogo/src/domain"
)

// TinyTest provides a simple testing API.
// Philosophy: Write tests like you write endpoints
type TinyTest struct {
	routeRegistry *application.RouteRegistry
}

// NewTinyTest creates a new test instance.
// Users: api := testing.NewTinyTest()
func NewTinyTest() *TinyTest {
	return &TinyTest{
		routeRegistry: application.NewRouteRegistry(),
	}
}

// POST registers a POST handler for testing.
// API Fluent: Same API as app.POST()
func (t *TinyTest) POST(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *TinyTest {
	t.registerRoute("POST", path, handler, opts...)
	return t
}

// GET registers a GET handler for testing.
func (t *TinyTest) GET(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *TinyTest {
	t.registerRoute("GET", path, handler, opts...)
	return t
}

// PUT registers a PUT handler for testing.
func (t *TinyTest) PUT(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *TinyTest {
	t.registerRoute("PUT", path, handler, opts...)
	return t
}

// DELETE registers a DELETE handler for testing.
func (t *TinyTest) DELETE(path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) *TinyTest {
	t.registerRoute("DELETE", path, handler, opts...)
	return t
}

// registerRoute registers a route for testing.
func (t *TinyTest) registerRoute(method, path string, handler domain.HandlerFunc, opts ...domain.RouteOptions) {
	var merged domain.RouteOptions
	for _, opt := range opts {
		if opt.Body != nil {
			merged.Body = opt.Body
		}
		if opt.Response != nil {
			merged.Response = opt.Response
		}
		if opt.Summary != "" {
			merged.Summary = opt.Summary
		}
		if len(opt.Tags) > 0 {
			merged.Tags = opt.Tags
		}
	}

	_ = t.routeRegistry.Register(method, path, handler, merged)
}

// ExpectSuccess makes a request and expects success.
// Returns test result with status code and body
func (t *TinyTest) ExpectSuccess(method, path string, body interface{}) *TestResult {
	return t.makeRequest(method, path, body, false)
}

// ExpectError makes a request and expects error.
// Returns test result with error details
func (t *TinyTest) ExpectError(method, path string, body interface{}) *TestResult {
	return t.makeRequest(method, path, body, true)
}

// makeRequest makes an HTTP request against registered routes.
func (t *TinyTest) makeRequest(method, path string, body interface{}, expectError bool) *TestResult {
	// Marshal body if provided
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return &TestResult{Error: err}
		}
	}

	// Create HTTP request
	req := httptest.NewRequest(method, path, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rec := httptest.NewRecorder()

	// Find route
	route := t.routeRegistry.Find(method, path)
	if route == nil {
		return &TestResult{StatusCode: 404, Error: domain.NewHTTPException(404, "route not found")}
	}

	// Create context
	ctx := &domain.Context{
		Params:      make(map[string]string),
		QueryParams: make(map[string]string),
		Headers:     make(map[string]string),
		StatusCode:  200,
	}

	// Bind query params
	for key, values := range req.URL.Query() {
		if len(values) > 0 {
			ctx.QueryParams[key] = values[0]
		}
	}

	// Bind headers
	for key, values := range req.Header {
		if len(values) > 0 {
			ctx.Headers[key] = values[0]
		}
	}

	// Store request
	ctx.Request = req
	ctx.Response = rec

	// Mock binder for tests
	ctx.Binder = nil // Tests don't use real binding

	// Call handler
	err := route.Handler(ctx)
	if err != nil {
		result := &TestResult{
			StatusCode: ctx.StatusCode,
			Error:      err,
		}
		
		// If expecting error, this is success
		if expectError {
			result.Success = true
		}
		
		return result
	}

	// Parse response
	var responseBody interface{}
	if len(rec.Body.Bytes()) > 0 {
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
	}

	// Determine success based on status code
	success := rec.Code >= 200 && rec.Code < 300

	return &TestResult{
		StatusCode: rec.Code,
		Body:       responseBody,
		Success:    success && !expectError,
	}
}

// TestResult represents the result of a test request.
type TestResult struct {
	StatusCode int
	Body       interface{}
	Error      error
	Success    bool
}

// Close cleans up test resources.
func (t *TinyTest) Close() {
	// Nothing to cleanup
}

