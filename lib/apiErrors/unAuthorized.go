package apierrors

import "Proyect-Y/common/errors"

type NotAuthorizedError struct {
	*errors.ProyectYError
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (no *NotAuthorizedError) Error() string {
  return no.Message
}
