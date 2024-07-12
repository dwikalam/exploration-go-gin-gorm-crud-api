package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"tutorial1.go.emp10.com/controllers"
	"tutorial1.go.emp10.com/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostsIndex)
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
