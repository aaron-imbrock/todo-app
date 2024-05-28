package main

import (
    "log"
	"net/http"
	"todo-app/commons/sqlite"
	"todo-app/features/todo"
    "github.com/gorilla/mux"
)

func main() {
	sqlite.InitDB()
	defer sqlite.DB.Close()

    r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.HandleFunc("/todos", todo.HandleGetTodos).Methods("GET")
	r.HandleFunc("/todos", todo.HandleCreateTodo).Methods("POST")
	r.HandleFunc("/todos/completion", todo.HandleUpdateTodoCompletion).Methods("PUT")
	r.HandleFunc("/todos", todo.HandleDeleteTodoByID).Methods("DELETE")

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
