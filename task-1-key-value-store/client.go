package main

import (
	"encoding/json"
	"os"
)

func getKey(path string, key string) (string, bool, error) {
	fileContentMap := map[string]string{}
	var fileContent, fileReadError = os.ReadFile(path)
	if fileReadError != nil {
		return "", not_found, fileReadError
	}
	marshallError := json.Unmarshal(fileContent, &fileContentMap)
	if marshallError != nil {
		return "", not_found, marshallError
	}
	if value, ok := fileContentMap[key]; !ok {
		return value, not_found, nil
	}
	return fileContentMap[key], found, nil
}

func setKey(path string, key string, value string) error {
	fileContentMap := map[string]string{}
	var fileContent, fileReadError = os.ReadFile(path)
	if fileReadError != nil {
		return fileReadError
	}
	unmarshallError := json.Unmarshal(fileContent, &fileContentMap)
	if unmarshallError != nil {
		return unmarshallError
	}
	fileContentMap[key] = value
	json, marshalError := json.Marshal(fileContentMap)
	if marshalError != nil {
		return marshalError
	}
	writeError := os.WriteFile(path, json, 0666)
	return writeError
}

const not_found = false
const found = true
