package core

import (
	"os"

	"sigs.k8s.io/yaml"
)

func YAMLToJSONFromFile(filenName string) (result []byte, err error) {
	var yamlBytes []byte
	yamlBytes, err = os.ReadFile(filenName)
	if err != nil {
		return
	}

	return yaml.YAMLToJSON(yamlBytes)
}

func JSONToYAMLFromFile(fileName string) (result []byte, err error) {
	var jsonBytes []byte
	jsonBytes, err = os.ReadFile(fileName)
	if err != nil {
		return
	}

	return yaml.JSONToYAML(jsonBytes)
}

func ReadJSONFile(fileName string) (result []byte, err error) {
	result, err = os.ReadFile(fileName)
	if err != nil {
		return result, err
	}
	return
}
