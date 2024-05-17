package controllers

import (
	"net/http"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	services "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Controller

	userService  services.UserService
	loginHandler gin.HandlerFunc
}

func CreateAuthController(userService services.UserService, loginHandler gin.HandlerFunc) *AuthController {
	return &AuthController{
		userService:  userService,
		loginHandler: loginHandler,
	}
}

func (con AuthController) Configure(r *gin.Engine) {
	g := r.Group("/api/v1/auth")
	{
		g.POST("/register", con.Register)
		g.POST("/login", con.Login)
	}
}

// UserRegister			godoc
// @Summary				Registers the user
// @Description			Checks the user data and adds it to the repo
// @Param				details body dto.PostUser true "Register details"
// @Tags				Auth
// @Success				200
// @Router				/auth/register [post]
func (con AuthController) Register(c *gin.Context) {
	var newUser dto.PostUser

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := con.userService.Register(&newUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusCreated)
}

// UserLogin			godoc
// @Summary				Logs in the user
// @Description			Checks the user data and returns a jwt token on correct Login
// @Param				details body dto.PostUser true "Login details"
// @Tags				Auth
// @Success				200 {object} services.LoginResult
// @Router				/auth/login [post]
func (con AuthController) Login(c *gin.Context) {
	con.loginHandler(c)
}
