package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	StatusCode  int    `json:"statusCode"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (r ResponseError) Error() string {

	return fmt.Sprintf("description: %s ,  code: %s", r.Description, r.Code)
}

func SetResponseError(statusCode int, code string, err *error) ResponseError {
	v := mapMessageError(code)
	validationErrors := (*err).(validator.ValidationErrors)

	if *err != nil && validationErrors != nil {
		for _, e := range validationErrors {
			v = mapMessageErrorValidator(e)
			break
		}
	} else if *err != nil {
		// for validator
		v = (*err).Error()
		code = "400"
	}

	return ResponseError{
		StatusCode:  statusCode,
		Code:        code,
		Description: v,
	}
}

func SetResponseData(metaData interface{}) map[string]any {

	value := map[string]any{"code": "0000", "description": "Success", "metaData": metaData}

	return value
}

func mapMessageError(code string) string {
	switch code {
	case "0000":
		return "Success"
	case "500":
		return "Internal Server Error"
	default:
		return "Unknown error"
	}
}

func mapMessageErrorValidator(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("This %s is required", err.Field())
	case "lte":
		return fmt.Sprintf("Should be less than %s", err.Field())
	case "gte":
		return fmt.Sprintf("Should be greater than %s", err.Field())
	}
	return "Unknown error"
}
