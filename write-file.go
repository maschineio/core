package core

import (
	"fmt"
	"os"
)

func WriteFile(fileName string, data []byte) error {
	// Create a new file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file", err)
		return err
	}
	defer file.Close()

	// Write the byte stream to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("write file", err)
		return err
	}
	return nil
}
