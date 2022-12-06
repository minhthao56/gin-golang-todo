package routers

import (
	"github.com/gin-gonic/gin"
)

type ControllersInterface interface {
	GetManyTodo(c *gin.Context)
	AddTodo(c *gin.Context)
	EditTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type TodoRoutersStruct struct {
	R                    *gin.Engine
	ControllersInterface ControllersInterface
}

func (r *TodoRoutersStruct) TodoRouters() {
	r.R.GET("/", r.ControllersInterface.GetManyTodo)
	r.R.POST("/add", r.ControllersInterface.AddTodo)
	r.R.PATCH("/:id", r.ControllersInterface.EditTodo)
	r.R.DELETE("/:id", r.ControllersInterface.DeleteTodo)
}
