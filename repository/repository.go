package repository

import (
	todo "TODO"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoStr interface {
	Create(listId int, item todo.TodoStr) (int, error)
	GetAll(userId, listId int) ([]todo.TodoStr, error)
	GetById(userId, strId int) (todo.TodoStr, error)
	Delete(userId, strId int) error
	UpdateStr(userId, strId int, input todo.UpdateStrInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoStr
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		TodoList:      NewTodoListDB(db),
		TodoStr:       NewTodoStrDB(db),
	}
}
