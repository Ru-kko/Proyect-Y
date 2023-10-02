package security

import "fmt"

type JWTError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (jw *JWTError) Error() string {
	return fmt.Sprintf("%s: %s", jw.Type, jw.Message)
}
