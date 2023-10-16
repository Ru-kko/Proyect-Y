package service

import (
	"Proyect-Y/gateway/internal/util"
	"Proyect-Y/typo"
	"encoding/json"
	"net/http"
)

func GetAuthentication(token string) (*typo.Auth, error) {
	if token == "" {
		return &typo.Auth{
			Authenticated: false,
		}, nil
	}

	authService := util.GetEnv().AUTH_ADDRESS

	req, err := http.NewRequest(http.MethodGet, authService+"/@me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", token)
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data interface{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	authData, ok := data.(typo.Auth)
	if !ok {
		return &typo.Auth{
			Authenticated: false,
		}, nil
	}

	return &authData, nil
}
