package main

import (
	"gin-golang/controllers"
	"gin-golang/database"
	"gin-golang/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := &routers.TodoRoutersStruct{
		R: gin.Default(),
		ControllersInterface: &controllers.ControllersStruct{
			DataBaseInterface: &database.DataBaseModelStruct{},
		},
	}

	r.TodoRouters()
	r.R.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
