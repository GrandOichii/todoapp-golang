package repositories

import (
	"context"

	"github.com/GrandOichii/todoapp-golang/api/config"
	"github.com/GrandOichii/todoapp-golang/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryImpl struct {
	TaskRepository

	config   *config.Configuration
	dbClient *mongo.Client
}

func CreateTaskRepositoryImpl(client *mongo.Client, config *config.Configuration) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{
		dbClient: client,
		config:   config,
	}
}

func (repo TaskRepositoryImpl) collection() *mongo.Collection {
	return repo.dbClient.Database(repo.config.Db.DbName).Collection(repo.config.Db.TaskCollection.Name)
}

func (repo TaskRepositoryImpl) FindAll() []*models.Task {
	var collection = repo.collection()
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var result []*models.Task
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (repo TaskRepositoryImpl) FindByOwnerId(ownerId string) []*models.Task {
	var collection = repo.collection()
	cursor, err := collection.Find(context.TODO(), bson.D{
		{Key: "ownerid", Value: ownerId},
	})
	if err != nil {
		panic(err)
	}
	var result []*models.Task
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (repo *TaskRepositoryImpl) Save(task *models.Task) bool {
	collection := repo.collection()
	insert, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		panic(err)
	}
	task.Id = insert.InsertedID.(primitive.ObjectID).Hex()
	return true
}

func (repo TaskRepositoryImpl) FindById(id string) *models.Task {
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	collection := repo.collection()

	find := collection.FindOne(context.TODO(), bson.D{
		{Key: "_id", Value: taskId},
	})

	err = find.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	var result models.Task
	err = find.Decode(&result)
	if err != nil {
		panic(err)
	}

	return &result
}

func (repo TaskRepositoryImpl) UpdateById(id string, updateF func(*models.Task) *models.Task) *models.Task {
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	collection := repo.collection()

	task := repo.FindById(id)
	if task == nil {
		return nil
	}

	newTask := updateF(task)
	newTask.Id = ""

	replace := collection.FindOneAndReplace(context.TODO(), bson.D{
		{Key: "_id", Value: taskId},
	}, newTask)

	err = replace.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	newTask.Id = id
	return newTask
}

func (repo *TaskRepositoryImpl) Remove(id string) bool {
	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}

	collection := repo.collection()

	delete := collection.FindOneAndDelete(context.TODO(), bson.D{
		{Key: "_id", Value: taskId},
	})

	err = delete.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		panic(err)
	}

	return true
}
