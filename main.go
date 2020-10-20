package main

import (
	"josmellbu/routers"
)

func main() {

	r := routers.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}