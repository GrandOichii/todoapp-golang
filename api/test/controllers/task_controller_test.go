package controllers_test

import (
	"errors"
	"testing"

	"github.com/GrandOichii/todoapp-golang/api/controllers"
	dto "github.com/GrandOichii/todoapp-golang/api/dto/task"
	taskservices "github.com/GrandOichii/todoapp-golang/api/services/task"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func createTaskController(taskService taskservices.TaskService) *controllers.TaskController {
	return controllers.CreateTaskController(
		taskService,
		// nil,
		func(*gin.Context) {},
		func(s string, ctx *gin.Context) (string, error) {
			return "userId", nil
		},
	)
}

func Test_ShouldFetchAll(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("GetAll", mock.Anything).Return([]*dto.GetTask{})

	c, w := createTestContext(nil)

	// act
	controller.All(c)

	// assert
	assert.Equal(t, w.Code, 200)
}

func Test_ShouldCreate(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("Add", mock.Anything, mock.Anything).Return(&dto.GetTask{}, nil)

	c, w := createTestContext(dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	})

	// act
	controller.Create(c)

	// assert
	assert.Equal(t, w.Code, 201)
}

func Test_ShouldNotCreate(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("Add", mock.Anything, mock.Anything).Return(nil, errors.New(""))

	c, w := createTestContext(dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	})

	// act
	controller.Create(c)

	// assert
	assert.Equal(t, w.Code, 400)
}

func Test_ShouldNotCreateBadData(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)

	c, w := createTestContext([]string{"first", "second"})

	// act
	controller.Create(c)

	// assert
	assert.Equal(t, w.Code, 400)
}

func Test_ShouldFetchById(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("GetById", mock.Anything, mock.Anything).Return(&dto.GetTask{}, nil)

	c, w := createTestContext(nil)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: "taskId",
	})

	// act
	controller.ById(c)

	// assert
	assert.Equal(t, w.Code, 200)
}

func Test_ShouldNotFetchById(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("GetById", mock.Anything, mock.Anything).Return(nil, errors.New(""))

	c, w := createTestContext(nil)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: "taskId",
	})

	// act
	controller.ById(c)

	// assert
	assert.Equal(t, w.Code, 404)
}

func Test_ShouldNotFetchByIdWithoutId(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("GetById", mock.Anything, mock.Anything).Return(nil, errors.New(""))

	c, w := createTestContext(nil)

	// act
	controller.ById(c)

	// assert
	assert.Equal(t, w.Code, 404)
}

func Test_ShouldToggleCompleted(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("ToggleCompleted", mock.Anything, mock.Anything).Return(&dto.GetTask{}, nil)

	c, w := createTestContext(nil)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: "taskId",
	})

	// act
	controller.ToggleCompleted(c)

	// assert
	assert.Equal(t, w.Code, 200)
}

func Test_ShouldNotToggleCompleted(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("ToggleCompleted", mock.Anything, mock.Anything).Return(nil, errors.New(""))

	c, w := createTestContext(nil)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: "taskId",
	})

	// act
	controller.ToggleCompleted(c)

	// assert
	assert.Equal(t, w.Code, 404)
}

func Test_ShouldNotToggleCompletedWithoutId(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("ToggleCompleted", mock.Anything, mock.Anything).Return(nil, errors.New(""))

	c, w := createTestContext(nil)

	// act
	controller.ToggleCompleted(c)

	// assert
	assert.Equal(t, w.Code, 404)
}

func Test_ShouldDelete(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("Delete", mock.Anything, mock.Anything).Return(nil)

	c, w := createTestContext(nil)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: "taskId",
	})

	// act
	controller.Delete(c)

	// assert
	assert.Equal(t, w.Code, 200)
}

func Test_ShouldNotDelete(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("Delete", mock.Anything, mock.Anything).Return(errors.New(""))

	c, w := createTestContext(nil)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: "taskId",
	})

	// act
	controller.Delete(c)

	// assert
	assert.Equal(t, w.Code, 404)
}

func Test_ShouldNotDeleteWithoutId(t *testing.T) {
	// arrange
	taskService := createMockTaskService()
	controller := createTaskController(taskService)
	taskService.On("Delete", mock.Anything, mock.Anything).Return(errors.New(""))

	c, w := createTestContext(nil)

	// act
	controller.Delete(c)

	// assert
	assert.Equal(t, w.Code, 404)
}
