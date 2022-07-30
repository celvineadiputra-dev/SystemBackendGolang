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

//func FormatValidationError(err error) []string {
//	var errors []string
//
//	for _, e := range err.(validator.ValidationErrors) {
//		errors = append(errors, e.Error())
//	}
//
//	return errors
//}

func customMessage(f validator.FieldError) string {
	if f.ActualTag() == "required" {
		return f.Field() + " is required"
	} else if f.ActualTag() == "eqfield" {
		return "Value not equal with " + f.Param()
	} else if f.ActualTag() == "min" {
		return "Minimum Character " + f.Param()
	}

	return f.ActualTag()
}

func FormatValidationError(err error) map[string]string {
	var verr validator.ValidationErrors

	errs := make(map[string]string)
	if errors.As(err, &verr) {
		for _, f := range verr {
			println(f.Error())
			println(f.ActualTag())
			println(f.Param())
			println(f.Field())
			err := f.ActualTag()
			if f.Param() != "" {
				err = fmt.Sprintf("%s=%s", err, f.Param())
			}
			println("XXX")
			println(f.ActualTag())
			println("XXX")
			errs[f.Field()] = customMessage(f)
		}
	}
	return errs
}
