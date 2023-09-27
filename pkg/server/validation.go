package server

import "fmt"

type Validatable interface {
	Validate() []error
}

type ValidationErrors struct {
	Errors []error
}

func (v *ValidationErrors) Error() (output string) {
	if len(v.Errors) == 0 {
		return
	}

	output = v.Errors[0].Error()

	if len(v.Errors) > 1 {
		output += fmt.Sprintf(" [+%d more error(s)]", len(v.Errors)-1)
	}

	return
}

func Validate(model any) *ValidationErrors {
	validatable, ok := model.(Validatable)
	if !ok {
		return nil
	}

	errs := validatable.Validate()
	if len(errs) == 0 {
		return nil
	}

	return &ValidationErrors{Errors: errs}
}
