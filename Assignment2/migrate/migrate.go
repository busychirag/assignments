package main

import "github.com/busychirag/assignments/tree/main/Assignment2/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	// initializers.DB.AutoMigrate()
}
