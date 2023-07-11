package main

import (
	"github.com/balqees/go-crud/initializers"
	"github.com/balqees/go-crud/routers"
	"github.com/balqees/go-crud/seed"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	seed.SeedUsers(initializers.DB)
	r := routers.SetupRouter()

	// Run the server
	r.Run()
}
