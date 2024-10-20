// internal/service/todo_service.go

package service

import (
	"time"

	"github.com/dotslashbit/todo-api/internal/model"
	"github.com/dotslashbit/todo-api/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(title string) (*model.Todo, error) {
	todo := &model.Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.repo.Create(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) GetTodo(id int64) (*model.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) ListTodos() ([]*model.Todo, error) {
	return s.repo.List()
}

func (s *TodoService) UpdateTodo(id int64, title string, completed bool) (*model.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	todo.Title = title
	todo.Completed = completed
	todo.UpdatedAt = time.Now()
	err = s.repo.Update(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) DeleteTodo(id int64) error {
	return s.repo.Delete(id)
}
