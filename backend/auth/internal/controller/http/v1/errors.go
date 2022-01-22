package v1

import "net/http"

type operationResult struct {
	Code        int
	Status      string
	Description string
}

var signUpValidationError = operationResult{
	Code:        http.StatusConflict,
	Status:      http.StatusText(http.StatusConflict),
	Description: "",
}
