package main

import (
	"fmt"
	"os"

	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	fmt.Println("Welcome to my Banking System")
	r := routes.SetupRouter()
	port := os.Getenv("PORT")
	err := r.Run(":" + port)

	if err != nil {
		fmt.Println("Error running the server", err)
		panic(err)
	}
	fmt.Println("Server running on port: ", port)
}
