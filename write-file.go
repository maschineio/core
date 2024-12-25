package core

import (
	"os"
)

func WriteFile(fileName string, data []byte) error {
	// Write the byte stream to the file using os.WriteFile
	err := os.WriteFile(fileName, data, 0640)
	if err != nil {
		return err
	}
	return nil
}
