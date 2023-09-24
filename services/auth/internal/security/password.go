package security

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(psw string) (string, error) {
	hashedPsw, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPsw), nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
