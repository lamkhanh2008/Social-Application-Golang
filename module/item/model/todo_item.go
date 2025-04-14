package model

import (
	"errors"
	"fmt"
	"social_todo/common"
	"strings"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
	EntityName            = "Todo_items"
)

type TodoItem struct {
	common.SQLModel
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:description;"`
	Status      string        `json:"status" gorm:"column:status;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
	UserId      int           `json:"user_id" gorm:"column:user_id;"`
}

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id;"`
	UserId      int    `json:"-" gorm:"column:user_id;"`
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

func (data *TodoItemCreation) Validate() error {
	title := strings.TrimSpace(data.Title)
	fmt.Println("title: ", title)
	if title != "" {
		return nil
	}

	return ErrTitleCannotBeEmpty
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
