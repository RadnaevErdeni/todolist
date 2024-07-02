package service

import (
	todo "TODO"
	"TODO/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoStr interface {
	Create(userId, listId int, item todo.TodoStr) (int, error)
	GetAll(userId, listId int) ([]todo.TodoStr, error)
	GetById(userId, strId int) (todo.TodoStr, error)
	Delete(userId, strId int) error
	UpdateStr(userId, listId int, input todo.UpdateStrInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoStr
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoStr:       NewTodoStrService(repos.TodoStr, repos.TodoList),
	}
}
