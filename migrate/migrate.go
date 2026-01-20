package main

import (
	"github.com/rhrashal/go-crud/initializers" // Adjust to your module name
	"github.com/rhrashal/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{}, &models.Product{}) // Add Product here
}
