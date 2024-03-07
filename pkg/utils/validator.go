package utils

import (
	"bike-rent-express/model/dto/json"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validated(s any) []json.ValidationField {
	var errors []json.ValidationField

	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())
		validate.RegisterValidation("format-date", validateDateFormat)
	}

	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldError := json.ValidationField{
				FieldName: err.Field(),
				Message:   getErrorMesssage(err.Tag()),
			}
			errors = append(errors, fieldError)
		}
	}
	return errors
}

func getErrorMesssage(tag string) string {
	messages := map[string]string{
		"required":    "field is required",
		"email":       "email is not valid",
		"string":      "field is not string",
		"number":      "field is not number",
		"format-date": "wrong date format",
	}

	for key, val := range messages {
		if tag == key {
			return val
		}
	}

	return ""
}

func validateDateFormat(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	dateFormat := regexp.MustCompile(`^(0[1-9]|1[0-2])-(0?[1-9]|[12][0-9]|3[01])-(\d{4})$`)
	return dateFormat.MatchString(date)
}
