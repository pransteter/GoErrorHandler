package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type ErrorData struct {
	Tag      string
	Message  string
	HttpCode int
}

const (
	GENERIC_ERROR          = 0
	INVALID_ARGUMENT_ERROR = 1
)

var errorMessages = map[int]string{
	GENERIC_ERROR:          "We have problems...",
	INVALID_ARGUMENT_ERROR: "The sent argument is invalid.",
}

var errorHttpCodes = map[int]int{
	GENERIC_ERROR:          400,
	INVALID_ARGUMENT_ERROR: 401,
}

var errorNames = map[int]string{
	GENERIC_ERROR:          "Generic",
	INVALID_ARGUMENT_ERROR: "Invalid Argument",
}

func ErrorMessage(errorCode int) string {
	return errorMessages[errorCode]
}

func ErrorHttpCode(errorCode int) int {
	return errorHttpCodes[errorCode]
}

func ErrorName(errorCode int) string {
	return errorNames[errorCode]
}

func ErrorHandler() {
	if err := recover(); err != nil {
		var errorData ErrorData

		if t := reflect.TypeOf(err); t.Kind().String() != "struct" || t.Name() != "ErrorData" {
			errorData.Tag = "[ Error Handler Error ]"
			errorData.Message = fmt.Sprintf("%s", err)
			errorData.HttpCode = 500
		}

		errorReflected := reflect.ValueOf(err)
		errorData = reflect.Indirect(
			errorReflected,
		).Interface().(ErrorData)

		fullErrorMessage := fmt.Sprintf("%s - %s", errorData.Tag, errorData.Message)

		log.Println(fullErrorMessage)
	}
}

func CreateHttpException(errorCode int, message ...string) ErrorData {
	getExceptionDataFromHttp := func(errorCode int) ErrorData {
		errorMessage := http.StatusText(errorCode)

		errorData := ErrorData{}
		errorData.Tag = fmt.Sprintf("[ %s Error ]", errorMessage)
		errorData.HttpCode = errorCode
		errorData.Message = errorMessage

		if len(message) > 0 {
			errorData.Message = message[0]
		}

		return errorData
	}

	errorData := ErrorData{}
	errorData = getExceptionDataFromHttp(errorCode)

	return errorData
}

func CreateException(errorCode int, message ...string) ErrorData {
	getExceptionDataFromList := func(errorCode int) ErrorData {
		errorData := ErrorData{}

		errorData.Tag = fmt.Sprintf("[ %s ]", ErrorName(errorCode))
		errorData.HttpCode = ErrorHttpCode(errorCode)
		errorData.Message = ErrorMessage(errorCode)

		if len(message) > 0 {
			errorData.Message = message[0]
		}

		return errorData
	}

	errorData := ErrorData{}
	errorData = getExceptionDataFromList(errorCode)

	return errorData
}
