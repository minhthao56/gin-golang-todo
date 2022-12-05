package routers

import (
	"gin-golang/controllers"

	"github.com/gin-gonic/gin"
)

type RoutersInterface interface {
	TodoRouters()
}

type TodoRoutersStruct struct {
	ControllersInterface controllers.ControllersInterface
}

var R = gin.Default()

func (r *TodoRoutersStruct) TodoRouters() {
	R.GET("/", r.ControllersInterface.GetManyTodo)
	R.POST("/add", r.ControllersInterface.AddTodo)
	R.PATCH("/:id", r.ControllersInterface.EditTodo)
	R.DELETE("/:id", r.ControllersInterface.DeleteTodo)
}
