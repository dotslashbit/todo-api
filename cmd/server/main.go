package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dotslashbit/todo-api/internal/handler"
	"github.com/dotslashbit/todo-api/internal/repository"
	"github.com/dotslashbit/todo-api/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	router := httprouter.New()
	router.POST("/todos", todoHandler.CreateTodo)
	router.GET("/todos/:id", todoHandler.GetTodo)
	router.GET("/todos", todoHandler.ListTodos)
	router.PUT("/todos/:id", todoHandler.UpdateTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// Start the server
	log.Printf("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
