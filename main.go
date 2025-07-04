package main

import (
	"backend-api/controllers"
	"backend-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//get database connection
	models.ConnectDatabase()

	//create route with method get
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	//create route for get all posts
	router.GET("/api/posts", controllers.FindPosts)
	//create route store posts
	router.POST("/api/posts", controllers.StorePost)
	//create route detail posts
	router.GET("/api/posts/:id", controllers.FindPostById)
	//create route update posts
	router.PUT("/api/posts/:id", controllers.UpdatePost)
	//create route delete posts
	router.DELETE("/api/posts/:id", controllers.DeletePosts)

	//start server with port 3000
	router.Run(":3000")
}
