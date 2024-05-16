package controllers

import (
	"net/http"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/task"
	services "github.com/GrandOichii/todoapp-golang/api/services/task"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Controller

	taskService services.TaskService
	auth        gin.HandlerFunc
}

func (con TaskController) Configure(r *gin.Engine) {
	g := r.Group("/api/v1/task")
	{
		g.Use(con.auth)
		g.GET("", con.all)
		g.POST("", con.create)
		g.DELETE(":id", con.delete)
		g.GET(":id", con.byId)
		g.PATCH(":id", con.toggleCompleted)
	}
}

func CreateTaskController(taskService services.TaskService, auth gin.HandlerFunc) *TaskController {
	return &TaskController{
		taskService: taskService,
		auth:        auth,
	}
}

// Find all tasks	godoc
// @Summary			Fetch all tasks
// @Description		Fetches all of the user's tasks
// @Tags			Tasks
// @Success			200 {object} []dto.GetTask
// @Router			/task [post]
func (con TaskController) all(c *gin.Context) {
	result := con.taskService.GetAll()
	c.IndentedJSON(http.StatusOK, result)
}

// TODO add swagger docs
func (con TaskController) create(c *gin.Context) {
	var newTask dto.CreateTask

	if err := c.BindJSON(&newTask); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result, err := con.taskService.Add(&newTask)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to create task: %s", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}

// TODO add swagger docs
func (con TaskController) byId(c *gin.Context) {
	id := c.Param("id")

	result, err := con.taskService.GetById(id)
	if err != nil {
		c.String(http.StatusNotFound, "failed to fetch task: %s", err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

// TODO add swagger docs
func (con TaskController) toggleCompleted(c *gin.Context) {
	id := c.Param("id")

	result, err := con.taskService.ToggleCompleted(id)
	if err != nil {
		c.String(http.StatusNotFound, "failed to patch task: %s", err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)

}

// TODO add swagger docs
func (con TaskController) delete(c *gin.Context) {
	id := c.Param("id")

	err := con.taskService.Delete(id)
	if err != nil {
		c.String(http.StatusNotFound, "failed to delete task: %s", err)
		return
	}

	c.Status(http.StatusOK)
}
