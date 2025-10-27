package application

import (
	"github.com/go-playground/validator/v10"
)

// SchemaValidator validates structs using struct tags.
// SOLID: Single Responsibility - only validates
// Adapter Pattern: Wraps go-playground/validator
type SchemaValidator struct {
	validator *validator.Validate
}

// NewSchemaValidator creates a new schema validator.
func NewSchemaValidator() *SchemaValidator {
	return &SchemaValidator{
		validator: validator.New(),
	}
}

// Validate validates a struct against its validation tags.
// Guard Clause: Fail fast if validation fails
func (s *SchemaValidator) Validate(data interface{}) error {
	// Guard clause: Validate first
	if err := s.validator.Struct(data); err != nil {
		return err // Fail fast
	}

	// Happy path: Validation passed
	return nil
}

// RegisterCustomValidator registers a custom validation function.
func (s *SchemaValidator) RegisterCustomValidator(tag string, fn validator.Func) error {
	return s.validator.RegisterValidation(tag, fn)
}

