package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Message string `json:"holeshit"`
}

func (r *ValidationError) Error() string {
	//indent, _ := json.MarshalIndent(r.Message, "", "  ")
	indent, _ := json.Marshal(r)
	return string(indent)
}

func Validation(i interface{}) error {
	validate := validator.New()

	validate.RegisterValidation("phone", PhoneNumberValidation)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		error := ""
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				error = fmt.Sprintf("field %s is required", err.Field())
			case "min":
				error = fmt.Sprintf("%s must have a minimum of %s characters", err.Field(), err.Param())
			case "max":
				error = fmt.Sprintf("%s must have a maximum of %s characters", err.Field(), err.Param())
			case "email":
				error = fmt.Sprintf("%s value does not match valid %s", err.Field(), err.Tag())
			case "phone":
				error = fmt.Sprintf("%s value does not match a valid phone number", err.Value())
			default:
				error = fmt.Sprintf("something wrong on %s", err.Field())
			}
		}

		return &ValidationError{
			Message: error,
		}

	}
	return nil
}

func PhoneNumberValidation(sl validator.FieldLevel) bool {
	value := sl.Field().String()
	re := regexp.MustCompile(`[+]{1}?\d{12}$`)
	return re.MatchString(value)
}
