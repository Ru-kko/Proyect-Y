package apierrors

import (
	"Proyect-Y/common/errors"
	"fmt"
)

type UserNotFound struct {
	*errors.ProyectYError
	Name    string `json:"name"`
	Message string `json:"message"`
	User    string `json:"user"`
}

func (us *UserNotFound) Error() string {
	us.Message = fmt.Sprintf("User with tag %s does not exists", us.Message)
	return us.Message
}

type ServiceNotFound struct {
	*errors.ProyectYError
	Name    string `json:"name"`
	Message string `json:"message"`
	Paht    string `json:"path"`
}

func (e ServiceNotFound) Error() string {
	return fmt.Sprintf("%s - %s", e.Message, e.Paht)
}
