package utils

import (
	"bike-rent-express/model/dto/json"

	"github.com/go-playground/validator/v10"
)

func Validated(s any) []json.ValidationField {
	var errors []json.ValidationField
	validate := validator.New(validator.WithRequiredStructEnabled())
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
		"required": "field is required",
		"email":    "email is not valid",
		"string":   "field is not string",
		"number":   "field is not number",
		"datetime": "wrong date format",
	}

	for key, val := range messages {
		if tag == key {
			return val
		}
	}

	return ""
}
