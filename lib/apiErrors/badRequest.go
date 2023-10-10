package apierrors

import (
	"Proyect-Y/common/errors"
	"fmt"
)

type BadRequest struct {
	*errors.ProyectYError
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (e BadRequest) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
