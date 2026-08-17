package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"_id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "01", Item: "Clean Room", Completed: true},
	{ID: "02", Item: "Read Book", Completed: false},
	{ID: "03", Item: "Workout", Completed: false},
	{ID: "04", Item: "Clean the car", Completed: true},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getByID(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getId(c *gin.Context) {
	id := c.Param("id")
	todo, err := getByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Todo Not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Todo not found / unable to update todo",
		})
		return
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getId)
	r.PATCH("/todos/:id", updateTodo)
	r.POST("/todos", addTodo)
	r.Run(":8000")
}
