package main

import (
    "log"

    "github.com/MarkTBSS/go-todo-api/config"
    "github.com/MarkTBSS/go-todo-api/models"
    "github.com/MarkTBSS/go-todo-api/routes"
)

func main() {
    config.ConnectDatabase()

    config.DB.AutoMigrate(&models.Todo{})
    log.Println("Database migrated successfully")

    r := routes.SetupRouter()

    port := ":8080"
    log.Printf("Server starting on port %s", port)
    r.Run(port)
}
