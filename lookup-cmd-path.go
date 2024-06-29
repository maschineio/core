package core

import (
	"fmt"
	"os/exec"
)

func LookupCmdPath(cmd string) (string, error) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", fmt.Errorf("cant get path for command '%s': %w", cmd, err)
	}
	return path, nil
}
