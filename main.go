package main

import (
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/routes"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnecToDB()
}

func main() {
	routes.Routes()
}
