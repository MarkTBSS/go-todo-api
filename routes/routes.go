package routes

import (
    "github.com/MarkTBSS/go-todo-api/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api/todos")
    {
        api.POST("", handlers.CreateTodo)
        api.GET("", handlers.GetTodos)
        api.GET("/:id", handlers.GetTodo)
        api.PUT("/:id", handlers.UpdateTodo)
        api.DELETE("/:id", handlers.DeleteTodo)
    }

    return r
}
