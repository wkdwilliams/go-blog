package validator

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
