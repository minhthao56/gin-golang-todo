package routers

import (
	"gin-golang/controllers"

	"github.com/gin-gonic/gin"
)

var R = gin.Default()

func TodoRouters() {
	R.GET("/", controllers.GetManyTodo)
	R.POST("/add", controllers.AddTodo)
	R.PATCH("/:id", controllers.EditTodo)
	R.DELETE("/:id", controllers.DeleteTodo)
}
