package repository

import (
	"github.com/Xploitable/todo-app"
	"github.com/jmoiron/sqlx"
)

const (
	USER_TABLE        = "users"
	TODO_LISTS_TABLE  = "todo_lists"
	USERS_LISTS_TABLE = "users_lists"
	TODO_ITEMS_TABLE  = "todo_items"
	LISTS_ITEMS_TABLE = "lists_items"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
