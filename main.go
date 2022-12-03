package main

import (
	"gin-golang/routers"
)

func main() {
	routers.TodoRouters()
	routers.R.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
