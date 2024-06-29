package core

import (
	"fmt"
	"os"
)

func GetEnvParameter(name string) (variable string, err error) {
	variable = os.Getenv(name)
	if variable == "" {
		return variable, fmt.Errorf("env variable %s not set", name)
	}
	return
}
