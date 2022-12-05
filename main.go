package main

import (
	"gin-golang/controllers"
	"gin-golang/database"
	"gin-golang/routers"
)

func main() {
	r := &routers.TodoRoutersStruct{
		ControllersInterface: &controllers.ControllersStruct{
			DataBaseModelInterface: &database.DataBaseModelStruct{},
		}}

	r.TodoRouters()
	routers.R.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
