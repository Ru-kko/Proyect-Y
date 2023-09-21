package util

import (
	apierrors "Proyect-Y/api-errors"
	"encoding/json"
	"io"
	"os"
)

type configFormat struct {
	Services map[string]string `json:"services"`
}

var conf *configFormat

func setConfig() error {
	env := GetEnv()

	file, err := os.Open(env.CONFIG)

	if err != nil {
		return err
	}
	defer file.Close()

	bitData, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	var res configFormat

	err = json.Unmarshal([]byte(bitData), &res)

	if err != nil {
		return err
	}

	conf = &res

	return nil
}

func GetService(name string) (string, error) {
	if conf == nil {
		setConfig()
	}
	host, exists := conf.Services[name]

	if !exists {
		return "", &apierrors.ServiceNotFound{Name: name}
	}

	return host, nil
}
