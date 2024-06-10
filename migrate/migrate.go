package main

import (
	"go-crud-example/initializers"
	"go-crud-example/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
