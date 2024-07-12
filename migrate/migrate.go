package main

import (
	"tutorial1.go.emp10.com/initializers"
	"tutorial1.go.emp10.com/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
