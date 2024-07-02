package service

import (
	todo "TODO"
	"TODO/repository"
)

type TodoStrService struct {
	repo     repository.TodoStr
	listRepo repository.TodoList
}

func NewTodoStrService(repo repository.TodoStr, listRepo repository.TodoList) *TodoStrService {
	return &TodoStrService{repo: repo, listRepo: listRepo}
}
func (s *TodoStrService) Create(userId, listId int, str todo.TodoStr) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, str)
}
func (s *TodoStrService) GetAll(userId, listId int) ([]todo.TodoStr, error) {
	return s.repo.GetAll(userId, listId)
}
func (s *TodoStrService) GetById(userId, strId int) (todo.TodoStr, error) {
	return s.repo.GetById(userId, strId)
}
func (s *TodoStrService) Delete(userId, strId int) error {
	return s.repo.Delete(userId, strId)
}

func (s *TodoStrService) UpdateStr(userId, listId int, input todo.UpdateStrInput) error {
	return s.repo.UpdateStr(userId, listId, input)
}
