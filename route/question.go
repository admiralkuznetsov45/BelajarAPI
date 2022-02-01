package route

import (
	"QuisAPI/controller"

	"github.com/labstack/echo"
)

func NewQuestion(e *echo.Echo) {
	e.POST("/question", controller.CreateQuestionController)
	e.GET("/question/:id", controller.GetQuestionByIDController)
	e.GET("/question", controller.GetAllQuestionsController)
	e.DELETE("/question/:id", controller.DeleteQuestionByIDController)
	e.PUT("/question/:id", controller.UpdateQuestionByIDController)

}
