package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/MarkTBSS/go-todo-api/config"
    "github.com/MarkTBSS/go-todo-api/models"
    "github.com/MarkTBSS/go-todo-api/routes"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func setupTestDB() {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&models.Todo{})
    config.DB = db
}

func TestCreateTodo(t *testing.T) {
    setupTestDB()
    router := routes.SetupRouter()

    body := map[string]string{
        "title":    "Buy groceries",
        "priority": "high",
    }
    jsonBody, _ := json.Marshal(body)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
    }

    var response models.Todo
    json.Unmarshal(w.Body.Bytes(), &response)
    if response.Title != "Buy groceries" {
        t.Errorf("Expected title 'Buy groceries', got '%s'", response.Title)
    }
    if response.Priority != "high" {
        t.Errorf("Expected priority 'high', got '%s'", response.Priority)
    }
}

func TestCreateTodoValidation(t *testing.T) {
    setupTestDB()
    router := routes.SetupRouter()

    // Missing required title
    body := map[string]string{"priority": "high"}
    jsonBody, _ := json.Marshal(body)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
    }
}

func TestGetTodos(t *testing.T) {
    setupTestDB()
    config.DB.Create(&models.Todo{Title: "Todo 1", Status: "pending", Priority: "medium"})
    config.DB.Create(&models.Todo{Title: "Todo 2", Status: "completed", Priority: "high"})
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/todos", nil)
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
    }

    var todos []models.Todo
    json.Unmarshal(w.Body.Bytes(), &todos)
    if len(todos) != 2 {
        t.Errorf("Expected 2 todos, got %d", len(todos))
    }
}

func TestGetTodosFilterByStatus(t *testing.T) {
    setupTestDB()
    config.DB.Create(&models.Todo{Title: "Todo 1", Status: "pending", Priority: "medium"})
    config.DB.Create(&models.Todo{Title: "Todo 2", Status: "completed", Priority: "high"})
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/todos?status=pending", nil)
    router.ServeHTTP(w, req)

    var todos []models.Todo
    json.Unmarshal(w.Body.Bytes(), &todos)
    if len(todos) != 1 {
        t.Errorf("Expected 1 todo, got %d", len(todos))
    }
}

func TestGetTodo(t *testing.T) {
    setupTestDB()
    config.DB.Create(&models.Todo{Title: "Test todo", Status: "pending", Priority: "medium"})
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/todos/1", nil)
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
    }
}

func TestGetTodoNotFound(t *testing.T) {
    setupTestDB()
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/todos/999", nil)
    router.ServeHTTP(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
    }
}

func TestUpdateTodo(t *testing.T) {
    setupTestDB()
    config.DB.Create(&models.Todo{Title: "Old title", Status: "pending", Priority: "medium"})
    router := routes.SetupRouter()

    body := map[string]string{"title": "New title", "status": "completed"}
    jsonBody, _ := json.Marshal(body)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/api/todos/1", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
    }

    var todo models.Todo
    json.Unmarshal(w.Body.Bytes(), &todo)
    if todo.Title != "New title" {
        t.Errorf("Expected title 'New title', got '%s'", todo.Title)
    }
    if todo.Status != "completed" {
        t.Errorf("Expected status 'completed', got '%s'", todo.Status)
    }
}

func TestUpdateTodoNotFound(t *testing.T) {
    setupTestDB()
    router := routes.SetupRouter()

    body := map[string]string{"title": "New title"}
    jsonBody, _ := json.Marshal(body)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/api/todos/999", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
    }
}

func TestDeleteTodo(t *testing.T) {
    setupTestDB()
    config.DB.Create(&models.Todo{Title: "Delete me", Status: "pending", Priority: "medium"})
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/api/todos/1", nil)
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
    }
}

func TestDeleteTodoNotFound(t *testing.T) {
    setupTestDB()
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/api/todos/999", nil)
    router.ServeHTTP(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
    }
}
