package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
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
	validate.RegisterValidation("date", DateValidation)

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
				error = fmt.Sprintf("vl: field %s is required", err.Field())
			case "min":
				error = fmt.Sprintf("vl: %s must have a minimum of %s characters", err.Field(), err.Param())
			case "max":
				error = fmt.Sprintf("vl: %s must have a maximum of %s characters", err.Field(), err.Param())
			case "email":
				error = fmt.Sprintf("vl: %s value does not match valid %s", err.Value(), err.Tag())
			case "phone":
				error = fmt.Sprintf("vl: %s value does not match a valid phone number", err.Value())
			case "uuid":
				error = fmt.Sprintf("vl: %s value does not match a valid id", err.Value())
			case "date":
				error = fmt.Sprintf("vl: %s value does not match a valid date", err.Value())
			default:
				error = fmt.Sprintf("vl: something wrong on %s", err.Field())
			}
		}

		return &ValidationError{Message: error}
	}
	return nil
}

func PhoneNumberValidation(sl validator.FieldLevel) bool {
	value := sl.Field().String()
	re := regexp.MustCompile(`[+]{1}?\d{12}$`)
	return re.MatchString(value)
}

func DateValidation(sl validator.FieldLevel) bool {
	value := sl.Field().String()
	if len(value) < 10 {
		return false
	}
	y, _ := strconv.Atoi(value[0:4])
	m, _ := strconv.Atoi(value[5:7])
	d, _ := strconv.Atoi(value[8:10])

	if d == 31 && (m == 4 || m == 6 || m == 9 || m == 11) {
		return false
	} else if m > 12 || (y == 0 || m == 0 || d == 0) {
		return false
	} else if d >= 30 && m == 2 {
		return false
	} else if m == 2 && d == 29 && !(y%4 == 0 && (y%100 != 0 || y%400 == 0)) {
		return false
	} else {
		return true
	}
}
