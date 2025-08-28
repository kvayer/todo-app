package repository

import (
	"github.com/jmoiron/sqlx"
	todoapp "github.com/kvayer/todo-app"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GetUser(username, password string) (todoapp.User, error)
}

type TodoList interface {
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(userId int) ([]todoapp.TodoList, error)
	GetById(userId, listId int) (todoapp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, id int, input todoapp.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todoapp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoapp.TodoItem, error)
	GetById(userId, itemId int) (todoapp.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoapp.UpdateItemInput) error
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
		TodoItem:      newTodoItemPostgres(db),
	}
}
