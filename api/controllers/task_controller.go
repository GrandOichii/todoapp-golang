package controllers

import (
	"net/http"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/task"
	"github.com/GrandOichii/todoapp-golang/api/middleware"
	services "github.com/GrandOichii/todoapp-golang/api/services/task"
	"github.com/GrandOichii/todoapp-golang/api/utility"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Controller

	taskService services.TaskService
	auth        gin.HandlerFunc
}

func (con TaskController) ConfigureApi(r *gin.Engine) {
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

func (con TaskController) ConfigureViews(r *gin.Engine) {
	g := r.Group("view")
	{
		g.Use(con.auth)
		g.GET("/tasks", con.GetTasks)
	}
}

func CreateTaskController(taskService services.TaskService, auth gin.HandlerFunc) *TaskController {
	return &TaskController{
		taskService: taskService,
		auth:        auth,
	}
}

// FindAllTasks		godoc
// @Summary			Fetch all tasks
// @Description		Fetches all of the user's tasks
// @Param			Authorization header string true "Authenticator"
// @Tags			Tasks
// @Success			200 {object} []dto.GetTask
// @Router			/task [get]
func (con TaskController) all(c *gin.Context) {
	userId, err := utility.Extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	result := con.taskService.GetAll(userId)
	c.IndentedJSON(http.StatusOK, result)
}

// CreateTask		godoc
// @Summary			Creates a task
// @Description		Creates a new user task
// @Param			Authorization header string true "Authenticator"
// @Param			task body dto.CreateTask true "new task data"
// @Tags			Tasks
// @Success			201 {object} dto.GetTask
// @Router			/task [post]
func (con TaskController) create(c *gin.Context) {
	userId, err := utility.Extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var newTask dto.CreateTask

	if err := c.BindJSON(&newTask); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result, err := con.taskService.Add(userId, &newTask)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}

// FindById			godoc
// @Summary			Find a task by id
// @Description		Finds a task by it's task id
// @Param			Authorization header string true "Authenticator"
// @Param   		taskId path string true "Task ID"
// @Tags			Tasks
// @Success			200 {object} dto.GetTask
// @Router			/task/{taskId} [get]
func (con TaskController) byId(c *gin.Context) {
	userId, err := utility.Extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	id := c.Param("id")

	result, err := con.taskService.GetById(userId, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

// ToggleComplete	godoc
// @Summary			Toggle complete status
// @Description		Toggles the task's complete status
// @Param			Authorization header string true "Authenticator"
// @Param   		taskId path string true "Task ID"
// @Tags			Tasks
// @Success			200 {object} dto.GetTask
// @Router			/task/{taskId} [patch]
func (con TaskController) toggleCompleted(c *gin.Context) {
	userId, err := utility.Extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	id := c.Param("id")

	result, err := con.taskService.ToggleCompleted(userId, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

// Delete			godoc
// @Summary			Delete task
// @Description		Deletes the task
// @Param			Authorization header string true "Authenticator"
// @Param   		taskId path string true "Task ID"
// @Tags			Tasks
// @Success			200
// @Router			/task/{taskId} [delete]
func (con TaskController) delete(c *gin.Context) {
	userId, err := utility.Extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	id := c.Param("id")

	err = con.taskService.Delete(userId, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
}

func (con *TaskController) GetTasks(c *gin.Context) {
	// c.JSON(http.StatusOK, "amogus")
	// return
	// con.auth(c)
	// status := c.Writer.Status()
	// if status != 401 {
	// 	c.HTML(status, "login", nil)
	// 	return
	// }

	// c.HTML(sta)
	userId, err := utility.Extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	tasks := con.taskService.GetAll(userId)
	c.HTML(http.StatusOK, "tasks", gin.H{
		"tasks": tasks,
	})
}
