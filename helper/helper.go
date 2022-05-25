package helper

import "gopkg.in/go-playground/validator.v9"

func ValidateStruct(user interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err3 := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err3.StructNamespace()
			element.Tag = err3.Tag()
			element.Value = err3.Value()
			errors = append(errors, &element)
		}
	}
	return errors
}
