package config

import (
	"context"
	"encoding/json"
	"os"

	"github.com/sethvargo/go-envconfig"
)

type CollectionConfiguration struct {
	Name string `json:"name" env:"NAME,required"`
}

type DbConfiguration struct {
	ConnectionUri  string                  `json:"connectionUri" env:"CONNECTION_URI,required"`
	DbName         string                  `json:"dbName" env:"NAME,required"`
	TaskCollection CollectionConfiguration `json:"taskCollection" env:",required,prefix=TASK_COLLECTION_"`
	UserCollection CollectionConfiguration `json:"userCollection" env:",required,prefix=USER_COLLECTION_"`
}

type Configuration struct {
	Port string          `json:"port" env:"PORT,required"`
	Db   DbConfiguration `json:"db" env:",required,prefix=DB_"`
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

func ReadEnvConfig() (*Configuration, error) {
	ctx := context.Background()

	var result Configuration
	if err := envconfig.Process(ctx, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
