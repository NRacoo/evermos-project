package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidatorResponse(err error) (validatorResponse []ValidationResponse) {
	var fieldErrors validator.ValidationErrors
	if errors.As(err, &fieldErrors) {

		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				validatorResponse = append(validatorResponse, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is required", err.Field()),
				})
			case "email":
				validatorResponse = append(validatorResponse, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is not a valid email", err.Field()),
				})
			default:
				ErrValidator, ok := ErrValidator[err.Tag()]
				if ok {
					count := strings.Count(ErrValidator, "%s")
					if count == 1 {
						validatorResponse = append(validatorResponse, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(ErrValidator, err.Field()),
						})
					} else {
						validatorResponse = append(validatorResponse, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(ErrValidator, err.Field(), err.Param()),
						})
					}
				} else {
					validatorResponse = append(validatorResponse, ValidationResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf("something wrong on %s : %s", err.Field(), err.Param()),
					})
				}

			}
		}
	}
	return validatorResponse
}

func WrapError(err error) error {
	logrus.Errorf("error : %s", err)
	return err
}
