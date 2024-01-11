package main

import (
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnecToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
