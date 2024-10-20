// internal/handler/todo_handler.go

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dotslashbit/todo-api/internal/service"
	"github.com/julienschmidt/httprouter"
)

type TodoHandler struct {
    todoService *service.TodoService
}

func NewTodoHandler(todoService *service.TodoService) *TodoHandler {
    return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var input struct {
        Title string `json:"title"`
    }
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    todo, err := h.todoService.CreateTodo(input.Title)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
    if err != nil {
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    todo, err := h.todoService.GetTodo(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    todos, err := h.todoService.ListTodos()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
    if err != nil {
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    var input struct {
        Title     string `json:"title"`
        Completed bool   `json:"completed"`
    }
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    todo, err := h.todoService.UpdateTodo(id, input.Title, input.Completed)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
    if err != nil {
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    err = h.todoService.DeleteTodo(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
