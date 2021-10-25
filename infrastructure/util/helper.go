package util

import (
	"encoding/json"
	"example/application/apperror"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Trans ut.Translator

// MustJSON is converter from interface{} to string
// Warning! this function will always assume the convertion is success
// if you are not sure the convertion is always succeed then use ToJSON
func MustJSON(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

// GetValidationErrorMessage is extractor error message from binding http request data
func GetValidationErrorMessage(err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return apperror.FailUnmarshalResponseBodyError
	}
	var errorMessages []string
	for _, e := range errs {
		errorMessages = append(errorMessages, e.Translate(Trans))
	}
	errorMessage := strings.Join(errorMessages, " \n")
	return apperror.ERR400.Var(errorMessage)
}
