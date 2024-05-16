package controllers

import (
	"net/http"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	services "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Controller

	userService services.UserService
}

func CreateAuthController(userService services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (con AuthController) Configure(r *gin.Engine) {
	// TODO
	g := r.Group("/api/v1/auth")
	{
		g.POST("/register", con.register)
		g.POST("/login", con.login)
	}
}

// UserRegister			godoc
// @Summary				Registers the user
// @Description			Checks the user data and adds it to the repo
// @Param				details body dto.PostUser true "register details"
// @Tags				Auth
// @Success				200
// @Router				/auth/register [post]
func (con AuthController) register(c *gin.Context) {
	var newUser dto.PostUser

	if err := c.BindJSON(&newUser); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err := con.userService.Register(&newUser)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to register: %s", err)
		return
	}

	c.Status(http.StatusCreated)
}

// UserLogin			godoc
// @Summary				Logs in the user
// @Description			Checks the user data and returns a jwt token on correct login
// @Param				details body dto.PostUser true "login details"
// @Tags				Auth
// @Success				200 {object} services.LoginResult
// @Router				/auth/login [post]
func (con AuthController) login(c *gin.Context) {
	var userData dto.PostUser

	if err := c.BindJSON(&userData); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	login, err := con.userService.Login(&userData)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to login: %s", err)
		return
	}

	c.IndentedJSON(http.StatusOK, login)
}
