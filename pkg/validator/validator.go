package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	mainValidation "github.com/go-playground/validator/v10"
)

func Validate(s any) error {
	validator := mainValidation.New(mainValidation.WithRequiredStructEnabled())

	if err := validator.Struct(s); err != nil {
		return parseErrors(s, err)
	}

	return nil
}

// 't' must be a pointer to a struct.
func parseErrors(t any, err error) error {
	tType, e := reflecType(t)
	if e != nil {
		return e
	}

	errors := make(errorMap)

	// I think the error will always be of type ValidationErrors...
	// But
	if validationErrors, ok := err.(mainValidation.ValidationErrors); ok {
		for _, valErr := range validationErrors {
			a := tagType(valErr, tType)

			if a == nil {
				continue
			}
			errMsg := strings.Split(valErr.Error(), "Error:")
			if len(errMsg) == 2 {
				parseErrorMessage(&errMsg[1])
				errors[*a] = strings.Split(valErr.Error(), "Error:")[1]
			} else {
				errors[*a] = valErr.Error()
			}

		}
	}

	if len(errors) == 0 {
		return nil
	}

	return ValidationErrors{
		Errors: errors,
	}
}

func parseErrorMessage(msg *string) {
	// Compile the regex
	re := regexp.MustCompile(`\b(for|'[^']+')\b`)

	// Replace the matches with an empty string
	result := re.ReplaceAllString(*msg, "")

	// Print the result
	fmt.Println(result)
}

func reflecType(t any) (reflect.Type, error) {
	tType := reflect.TypeOf(t)

	if tType == nil || tType.Kind() != reflect.Ptr || tType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("t must be a pointer to a struct")
	}

	return tType, nil
}

func tagType(fieldError mainValidation.FieldError, fieldType reflect.Type) *string {
	field, exists := fieldType.Elem().FieldByName(fieldError.Field())
	if !exists {
		return nil
	}

	if fieldTag, exists := field.Tag.Lookup("form"); exists {
		return &fieldTag
	} else if fieldTag, exists := field.Tag.Lookup("json"); exists {
		return &fieldTag
	}

	return nil
}
