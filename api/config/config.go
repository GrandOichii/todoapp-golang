package config

import (
	"encoding/json"
	"os"
)

type CollectionConfiguration struct {
	Name string `json:"name"`
}

type DbConfiguration struct {
	ConnectionUri  string                  `json:"connectionUri"`
	DbName         string                  `json:"dbName"`
	TaskCollection CollectionConfiguration `json:"taskCollection"`
	UserCollection CollectionConfiguration `json:"userCollection"`
}

type Configuration struct {
	Port string          `json:"port"`
	Db   DbConfiguration `json:"db"`
}

func ReadConfig(path string) (*Configuration, error) {
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
