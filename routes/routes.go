package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/srbmaury/go-todo/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		// v1.POST("/todos/", controllers.postTodo)
		v1.GET("/todos", controllers.getTodos)
		// v1.GET("/todos/:id", controllers.getTodo)
		// v1.PUT("/todos/:id", controllers.updateTodo)
		// v1.PATCH("/todos/:id", controllers.partiallyUpdateTodo)
		// v1.DELETE("/todos/:id", controllers.deleteTodo)

		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}

	return r
}
