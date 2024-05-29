package main

import (
	"github.com/gorilla/mux"
	"log"
	"time"
	"net/http"
	"todo-app/commons/sqlite"
	"todo-app/features/todo"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
	})
}

func main() {
	sqlite.InitDB()
	defer sqlite.DB.Close()

	r := mux.NewRouter()
	r.Use(LoggingMiddleware)

	r.HandleFunc("/", todo.HandleGetTodos).Methods("GET") 
	r.HandleFunc("/todos", todo.HandleGetTodos).Methods("GET")
	r.HandleFunc("/todos", todo.HandleCreateTodo).Methods("POST")
	r.HandleFunc("/todos/completion", todo.HandleUpdateTodoCompletion).Methods("PUT")
	r.HandleFunc("/todos", todo.HandleDeleteTodoByID).Methods("DELETE")
	r.HandleFunc("/add-test-todo", todo.HandleAddTestTodo).Methods("GET")

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
