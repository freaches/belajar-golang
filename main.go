package main

import (
	"micro-feature-pemilu/controllers"
	"micro-feature-pemilu/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnecToDB()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.PATCH("/post/:id", controllers.PostsUpdate)
	r.GET("/post", controllers.PostIndex)
	r.GET("/post/:id", controllers.PostShow)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.Run()
}
