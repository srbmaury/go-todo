package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srbmaury/go-todo/models"
)

// func postTodo(c *gin.Context) {
// 	var newTodo Todo

// 	if error := c.BindJSON(&newTodo); error != nil {
// 		return
// 	}

// 	todos = append(todos, newTodo)

// 	c.IndentedJSON(http.StatusCreated, newTodo)
// }

func getTodos(c *gin.Context) {
	var todos []models.Todo
	c.IndentedJSON(http.StatusOK, todos)
}

// func getTodo(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range todos {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
// }

// func updateTodo(c *gin.Context) {
// 	id := c.Param("id")

// 	var updateData todo

// 	if error := c.BindJSON(&updateData); error != nil {
// 		return
// 	}
// 	for i, todo := range todos {
// 		if todo.ID == id {
// 			todos[i] = updateData
// 			c.IndentedJSON(http.StatusOK, todos[i])
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
// }

// func partiallyUpdateTodo(c *gin.Context) {
// 	id := c.Param("id")

// 	var updatedTodo struct {
// 		Title       *string `json:"title"`
// 		Description *string `json:"description"`
// 	}

// 	if err := c.BindJSON(&updatedTodo); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for i, todo := range todos {
// 		if todo.ID == id {
// 			if updatedTodo.Title != nil {
// 				todos[i].Title = *updatedTodo.Title
// 			}
// 			if updatedTodo.Description != nil {
// 				todos[i].Description = *updatedTodo.Description
// 			}

// 			c.JSON(http.StatusOK, todos[i])
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
// }

// func RemoveIndex(s []todo, index int) []todo {
// 	return append(s[:index], s[index+1:]...)
// }
// func deleteTodo(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, todo := range todos {
// 		if todo.ID == id {
// 			todos = RemoveIndex(todos, i)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Todo not found"})
// }
