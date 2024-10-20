// internal/repository/todo_repository.go

package repository

import (
	"github.com/dotslashbit/todo-api/internal/model"
	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
    db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
    return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *model.Todo) error {
    query := `INSERT INTO todos (title, completed, created_at, updated_at)
              VALUES ($1, $2, $3, $4) RETURNING id`
    return r.db.QueryRow(query, todo.Title, todo.Completed, todo.CreatedAt, todo.UpdatedAt).Scan(&todo.ID)
}

func (r *TodoRepository) GetByID(id int64) (*model.Todo, error) {
    var todo model.Todo
    err := r.db.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
    if err != nil {
        return nil, err
    }
    return &todo, nil
}

func (r *TodoRepository) List() ([]*model.Todo, error) {
    var todos []*model.Todo
    err := r.db.Select(&todos, "SELECT * FROM todos ORDER BY created_at DESC")
    return todos, err
}

func (r *TodoRepository) Update(todo *model.Todo) error {
    query := `UPDATE todos SET title = $1, completed = $2, updated_at = $3 WHERE id = $4`
    _, err := r.db.Exec(query, todo.Title, todo.Completed, todo.UpdatedAt, todo.ID)
    return err
}

func (r *TodoRepository) Delete(id int64) error {
    _, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
    return err
}
