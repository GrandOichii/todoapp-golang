package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port string `json:"port"`
}

func readConfig(path string) (*Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	result := &Configuration{}
	err = decoder.Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
