package models

import "time"

type Todo struct {
    ID          uint       `json:"id" gorm:"primaryKey"`
    Title       string     `json:"title" gorm:"size:255;not null"`
    Description string     `json:"description" gorm:"type:text"`
    Status      string     `json:"status" gorm:"size:20;default:'pending'"`
    Priority    string     `json:"priority" gorm:"size:10;default:'medium'"`
    DueDate     *time.Time `json:"due_date"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateTodoInput struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
    Status      string `json:"status"`
    Priority    string `json:"priority"`
    DueDate     string `json:"due_date"`
}

type UpdateTodoInput struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
    Priority    string `json:"priority"`
    DueDate     string `json:"due_date"`
}
