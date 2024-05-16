package repositories

import (
	"context"

	"github.com/GrandOichii/todoapp-golang/api/config"
	"github.com/GrandOichii/todoapp-golang/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	UserRepository

	dbClient *mongo.Client
	config   *config.Configuration
}

func (repo UserRepositoryImpl) collection() *mongo.Collection {
	return repo.dbClient.Database(repo.config.Db.DbName).Collection(repo.config.Db.UserCollection.Name)
}

func (repo UserRepositoryImpl) FindByUsername(username string) *models.User {
	collection := repo.collection()

	find := collection.FindOne(context.TODO(), bson.D{
		{Key: "username", Value: username},
	})

	err := find.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}

	var result models.User
	err = find.Decode(&result)
	if err != nil {
		panic(err)
	}

	return &result
}
