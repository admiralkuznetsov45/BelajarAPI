package route

import (
	"QuisAPI/controller"

	"github.com/labstack/echo"
)

func NewUser(e *echo.Echo) {

	e.POST("/users/login", controller.LoginUserController)
	e.GET("/users", controller.GetAllUsersController)
	e.POST("/users/register", controller.CreateUserController)
}
