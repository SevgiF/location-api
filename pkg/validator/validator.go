package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var Validate *CustomValidator

type CustomValidator struct {
	Validator *validator.Validate
}

// InitValidator initializes and configures the CustomValidator
func InitValidator() *CustomValidator {
	validate := NewValidator()
	//add custom validations
	validate.RegisterCustomValidations()
	return validate
}

// NewValidator creates a basic CustomValidator
func NewValidator() *CustomValidator {
	validate := validator.New()
	return &CustomValidator{Validator: validate}
}

// RegisterCustomValidations is a method of CustomValidator and
// registers custom validation rules
func (cv *CustomValidator) RegisterCustomValidations() {
	cv.Validator.RegisterValidation("hexadecimal", cv.isHexadecimal)
}

// isHexadecimal is a special validation method that checks if a field has a hexadecimal color code (6 characters)
func (cv *CustomValidator) isHexadecimal(fl validator.FieldLevel) bool {
	hexadecimalColorRegex := regexp.MustCompile(`^[0-9a-fA-F]{6}$`)
	return hexadecimalColorRegex.MatchString(fl.Field().String())
}
