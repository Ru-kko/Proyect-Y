package apierrors

import "Proyect-Y/common/errors"

type InternalServerError struct {
	*errors.ProyectYError
	Name    string `json:"name"`
	Message string `json:"message,omitempty"`
	Code    int16  `json:"code"`
}

func (e InternalServerError) Error() string {
	return "Internal server error"
}
