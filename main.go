package main

import (
	"github.com/balqees/go-crud/initializers"
	"github.com/balqees/go-crud/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := routers.SetupRouter()

	// Run the server
	r.Run()
}
