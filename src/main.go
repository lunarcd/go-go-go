package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lunarcd/go-go-go/src/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{}

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/todos", GetTodos)
	r.GET("/todos/:id", GetTodoByID)
	r.POST("/todos", CreateTodo)
	r.PUT("/todos/:id", UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)

	r.Run(":8080")
}

// GetTodos godoc
// @Summary      Get all todos
// @Description  Returns list of all todos
// @Tags         Todos
// @Produce      json
// @Success      200  {array}   Todo
// @Router       /todos [get]
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
// @Summary      Get todo by ID
// @Description  Returns a single todo
// @Tags         Todos
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  Todo
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [get]
func GetTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, t := range todos {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// CreateTodo godoc
// @Summary      Create new todo
// @Description  Add a new todo item
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        todo  body      Todo  true  "New Todo"
// @Success      201   {object}  Todo
// @Failure      400   {object}  map[string]string
// @Router       /todos [post]
func CreateTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// UpdateTodo godoc
// @Summary      Update todo by ID
// @Description  Update an existing todo
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        id    path      int   true  "Todo ID"
// @Param        todo  body      Todo  true  "Updated Todo"
// @Success      200   {object}  Todo
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /todos/{id} [put]
func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTodo Todo
	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Done = updatedTodo.Done
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// DeleteTodo godoc
// @Summary      Delete todo by ID
// @Description  Remove a todo
// @Tags         Todos
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
