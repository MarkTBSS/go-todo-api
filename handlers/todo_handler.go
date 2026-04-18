package handlers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/MarkTBSS/go-todo-api/config"
    "github.com/MarkTBSS/go-todo-api/models"
    "github.com/gin-gonic/gin"
)

// Create
func CreateTodo(c *gin.Context) {
    var input models.CreateTodoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    todo := models.Todo{
        Title:       input.Title,
        Description: input.Description,
        Status:      input.Status,
        Priority:    input.Priority,
    }

    if todo.Status == "" {
        todo.Status = "pending"
    }
    if todo.Priority == "" {
        todo.Priority = "medium"
    }

    if input.DueDate != "" {
        dueDate, err := time.Parse("2006-01-02", input.DueDate)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date format. Use YYYY-MM-DD"})
            return
        }
        todo.DueDate = &dueDate
    }

    if err := config.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
        return
    }

    c.JSON(http.StatusCreated, todo)
}

// Read All
func GetTodos(c *gin.Context) {
    var todos []models.Todo
    query := config.DB.Model(&models.Todo{})

    if status := c.Query("status"); status != "" {
        query = query.Where("status = ?", status)
    }
    if priority := c.Query("priority"); priority != "" {
        query = query.Where("priority = ?", priority)
    }

    query.Order("created_at DESC").Find(&todos)
    c.JSON(http.StatusOK, todos)
}

// Read One
func GetTodo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var todo models.Todo
    if err := config.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

// Update
func UpdateTodo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var todo models.Todo
    if err := config.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    var input models.UpdateTodoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.Title != "" {
        todo.Title = input.Title
    }
    if input.Description != "" {
        todo.Description = input.Description
    }
    if input.Status != "" {
        todo.Status = input.Status
    }
    if input.Priority != "" {
        todo.Priority = input.Priority
    }
    if input.DueDate != "" {
        dueDate, err := time.Parse("2006-01-02", input.DueDate)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date format. Use YYYY-MM-DD"})
            return
        }
        todo.DueDate = &dueDate
    }

    config.DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

// Delete
func DeleteTodo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    result := config.DB.Delete(&models.Todo{}, id)
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
