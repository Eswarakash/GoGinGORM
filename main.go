package main

import (
	"github.com/Eswarakash/GoGinGORM/controllers"
	"github.com/Eswarakash/GoGinGORM/initializers"
	"github.com/Eswarakash/GoGinGORM/migrate"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.MigrateAuto()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.GetPong)
	r.POST("/Post", controllers.PostsCreate)
	r.GET("/Get", controllers.GetPosts)
	r.GET("/Get/:id", controllers.GetPost)
	r.PUT("/Update/:id", controllers.UpdatePost)
	r.DELETE("/Delete/:id", controllers.DeletePost)

	r.Run()
}
