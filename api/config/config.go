package config

import (
	"context"
	"encoding/json"
	"os"

	"github.com/sethvargo/go-envconfig"
)

// type CollectionConfiguration struct {
// 	Name string `json:"name"`
// }

// type DbConfiguration struct {
// 	ConnectionUri  string                  `json:"connectionUri"`
// 	DbName         string                  `json:"dbName" env:"NAME,required"`
// 	TaskCollection CollectionConfiguration `json:"taskCollection"`
// 	UserCollection CollectionConfiguration `json:"userCollection"`
// }

// type Configuration struct {
// 	Port string          `json:"port" env:"PORT,required"`
// 	Db   DbConfiguration `json:"db" env:"DB,required"`
// }

type Configuration struct {
	Port string `json:"port" env:"PORT,required"`
	Db   struct {
		ConnectionUri  string `json:"connectionUri" env:"DB_CONNECTION_URI,required"`
		DbName         string `json:"dbName" env:"DB_NAME,required"`
		TaskCollection struct {
			Name string `json:"name" env:"DB_TASK_COLLECTION_NAME,required"`
		} `json:"taskCollection"`
		UserCollection struct {
			Name string `json:"name" env:"DB_USER_COLLECTION_NAME,required"`
		} `json:"userCollection"`
	} `json:"db"`
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
