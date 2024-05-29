package todo

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "strconv"
)

var tmpl = template.Must(template.ParseFiles(filepath.Join(".", "templates", "index.html")))

func HandleGetTodos(w http.ResponseWriter, r *http.Request) {
    log.Println("HandleGetTodos called")

    todos, err := GetTodos()

    if err != nil {
        log.Printf("HandleGetTodos Error getting todos: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Printf("Rendering template with todos: %v", todos)
    err = tmpl.Execute(w, todos)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    _, err := CreateTodo(todo.Title)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func HandleUpdateTodoCompletion(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    completed := r.URL.Query().Get("completed") == "true"
    _, err = UpdateTodoCompletion(id, completed)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func HandleDeleteTodoByID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    _, err = DeleteTodoByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func HandleAddTestTodo(w http.ResponseWriter, r *http.Request) {
    _, err := CreateTodo("Test Todo Item")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write([]byte("Test Todo Added"))
}
