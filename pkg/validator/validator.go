package validator

import (
	"errors"
	"reflect"

	v "github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *v.Validate
}

func (v Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewValidator() Validator {
	return Validator{
		validator: v.New(),
	}
}

type errorMap map[string]string

type ValidationErrors struct {
	Errors errorMap `json:"errors"`
}

func (j ValidationErrors) Error() string {
	return "validation errors occurred"
}

func (j ValidationErrors) Is(err error) bool {
	_, ok := err.(ValidationErrors)
	return ok
}

// 't' must be a pointer to a struct. The 'lookup' must be either "form" or "json".
func ParseErrors(t any, err error, lookup string) error {
	// Ensure lookup is either "form" or "json"
	if lookup != "form" && lookup != "json" {
		return errors.New("lookup must be either 'form' or 'json'")
	}

	// Ensure 't' is a pointer to a struct
	tType := reflect.TypeOf(t)
	if tType == nil || tType.Kind() != reflect.Ptr || tType.Elem().Kind() != reflect.Struct {
		return errors.New("t must be a pointer to a struct")
	}

	errors := make(errorMap)

	if validationErrors, ok := err.(v.ValidationErrors); ok {
		for _, valErr := range validationErrors {
			field, exists := tType.Elem().FieldByName(valErr.Field())
			if !exists {
				continue
			}

			fieldTag, exists := field.Tag.Lookup(lookup)
			if !exists {
				fieldTag = valErr.Field()
			}

			errors[fieldTag] = valErr.Tag()
		}
	}

	if len(errors) == 0 {
		return nil
	}

	return ValidationErrors{
		Errors: errors,
	}
}
