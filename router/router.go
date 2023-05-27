package router

import (
	todosService "Demo/internal/todos/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/todos/", todosService.PostTodo)
		v1.GET("/todos", todosService.GetTodos)
		v1.GET("/todos/:id", todosService.GetTodo)
		v1.PUT("/todos/:id", todosService.UpdateTodo)
		v1.PATCH("/todos/:id", todosService.PartiallyUpdateTodo)
		v1.DELETE("/todos/:id", todosService.DeleteTodo)

		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}

	return r
}
