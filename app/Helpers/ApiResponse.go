package Helpers

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func ApiResponse(code int, message string, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Message: message,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) map[string]string {
	var verr validator.ValidationErrors

	errs := make(map[string]string)
	if errors.As(err, &verr) {

		for _, f := range verr {
			err := f.ActualTag()
			if f.Param() != "" {
				err = fmt.Sprintf("%s=%s", err, f.Param())
			}
			errs[f.Field()] = err
		}

	}
	return errs
}
