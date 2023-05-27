package todosService

import (
	todoModels "Demo/internal/todos/models"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var todos = []todoModels.Todo{}

func PostTodo(c *gin.Context) {
	var newTodo todoModels.Todo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo.ID = uuid.NewV4()
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, "Todo created")
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	parsedID := uuid.FromStringOrNil(id)

	if parsedID == uuid.Nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Invalid UUID"})
		return
	}

	for _, a := range todos {
		if a.ID == parsedID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var updateData todoModels.Todo

	parsedID := uuid.FromStringOrNil(id)
	if parsedID == uuid.Nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Invalid UUID"})
		return
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == parsedID {
			todos[i].Title = updateData.Title
			todos[i].Description = updateData.Description
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
}

func PartiallyUpdateTodo(c *gin.Context) {
	id := c.Param("id")

	parsedID := uuid.FromStringOrNil(id)
	if parsedID == uuid.Nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Invalid UUID"})
		return
	}

	var updatedTodo struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == parsedID {
			if updatedTodo.Title != nil {
				todos[i].Title = *updatedTodo.Title
			}
			if updatedTodo.Description != nil {
				todos[i].Description = *updatedTodo.Description
			}

			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func RemoveIndex(s []todoModels.Todo, index int) []todoModels.Todo {
	return append(s[:index], s[index+1:]...)
}
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	parsedID := uuid.FromStringOrNil(id)
	if parsedID == uuid.Nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Invalid UUID"})
		return
	}

	for i, todo := range todos {
		if todo.ID == parsedID {
			todos = RemoveIndex(todos, i)
			c.IndentedJSON(http.StatusOK, "Deleted Successfully")
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
}
