package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type ValidationError struct {
	Message string `json:"error"`
}

func (r *ValidationError) Error() string {
	//error, _ := json.MarshalIndent(r.Message, "", "  ")
	error, _ := json.Marshal(r)
	return string(error)
}

func Validation(i interface{}) error {
	validate := validator.New()

	validate.RegisterValidation("phone", PhoneNumberValidation)
	validate.RegisterValidation("array_uuid", ArrayUUIDValidation)
	validate.RegisterValidation("is_slice", IsSlice)
	validate.RegisterValidation("is_string_elem", IsStringElem)

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
			// case "uuid":
			// 	error = fmt.Sprintf("%s value does not match a valid %s's id", err.Value(), err.Field())
			case "array_uuid":
				error = fmt.Sprintf("one or more values do not match a valid %s id", err.Field())
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

func ArrayUUIDValidation(sl validator.FieldLevel) bool {
	value := sl.Field()
	re := regexp.MustCompile(`^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`)

	slice, ok := value.Interface().(pq.StringArray)
	if !ok {
		return false
	}
	for _, v := range slice {
		return re.MatchString(v)
	}
	return true
}

func IsSlice(fl validator.FieldLevel) bool {
	return fl.Top().Kind() == reflect.Slice
}

func IsStringElem(fl validator.FieldLevel) bool {
	return fl.Top().Type().Elem().Kind() == reflect.String
}
