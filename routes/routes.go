package routes

import (
	"micro-feature-pemilu/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.PATCH("/post/:id", controllers.PostsUpdate)
	r.GET("/post", controllers.PostIndex)
	r.GET("/post/:id", controllers.PostShow)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.POST("/user", controllers.TestsCreate)
	r.GET("/user", controllers.TestIndex)
	r.Run()
}
