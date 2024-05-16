package services_test

import (
	"testing"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/task"
	"github.com/GrandOichii/todoapp-golang/api/models"
	services "github.com/GrandOichii/todoapp-golang/api/services/task"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createTaskService(repo *MockTaskRepository) services.TaskService {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return services.CreateTaskServiceImpl(
		repo,
		validate,
	)
}

func Test_ShouldGetAll(t *testing.T) {
	// arrange
	repo := createTaskRepository()
	service := createTaskService(repo)
	repo.On("FindByOwnerId", mock.Anything).Return([]*models.Task{})

	// act
	tasks := service.GetAll("user")

	// assert
	assert.NotNil(t, tasks)
	assert.Len(t, tasks, 0)
}

func Test_ShouldCreate(t *testing.T) {
	// arrange
	repo := createTaskRepository()
	service := createTaskService(repo)
	repo.On("Save", mock.Anything).Return(true)

	// act
	tasks, err := service.Add("user", &dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	})

	// assert
	assert.NotNil(t, tasks)
	assert.Nil(t, err)
}

func Test_ShouldGetById(t *testing.T) {
	// arrange
	repo := createTaskRepository()
	service := createTaskService(repo)
	taskId := "taskid"
	userId := "userid"
	repo.On("FindById", mock.Anything).Return(&models.Task{
		OwnerId: userId,
		Id:      taskId,
	})

	// act
	task, err := service.GetById(userId, taskId)

	// assert
	assert.NotNil(t, task)
	assert.Nil(t, err)
}

func Test_ShouldNotGetById(t *testing.T) {
	// arrange
	repo := createTaskRepository()
	service := createTaskService(repo)
	taskId := "taskid"
	userId := "userid"

	repo.On("FindById", mock.Anything).Return(nil)

	// act
	task, err := service.GetById(userId, taskId)

	// assert
	assert.Nil(t, task)
	assert.NotNil(t, err)
}

func Test_ShouldNotGetByStolenId(t *testing.T) {
	// arrange
	repo := createTaskRepository()
	service := createTaskService(repo)
	taskId := "taskid"
	userId := "userid"
	repo.On("FindById", mock.Anything).Return(&models.Task{
		OwnerId: "otherUserId",
		Id:      taskId,
	})

	// act
	task, err := service.GetById(userId, taskId)

	// assert
	assert.Nil(t, task)
	assert.NotNil(t, err)
}

// TODO add ToggleCompleted tests
// TODO add Delete tests
