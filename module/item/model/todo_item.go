package model

import (
	"social_todo/common"
	"strings"
)

type TodoItem struct {
	common.SQLModel
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

type TodoItemUpdate struct {
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string { return "todo_items" }

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

func (data *TodoItemCreation) Validate() bool {
	title := strings.TrimSpace(data.Title)
	if title != "" {
		return true
	}

	return false
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}

func (data *TodoItemUpdate) Validate() bool {
	if strings.TrimSpace(data.Title) != "" {
		return true
	}

	return false
}
