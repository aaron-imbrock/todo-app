package todo

import (
    "database/sql"
    "log"
    "todo-app/commons/sqlite"
)

type Todo struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Completed   bool    `json:"completed"`
}

func GetTodos() ([]Todo, error) {
    rows, err := sqlite.DB.Query("SELECT id, title, completed FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    todos := []Todo{}
    for rows.Next() {
        var todo Todo
        var completed int
        if err := rows.Scan(&todo.ID, &todo.Title, &completed); err != nil {
            return nil, err
        }
        todo.Completed = completed == 1
        todos = append(todos, todo)
    }
    return todos, nil
}

func CreateTodo(title string) (sql.Result, error) {
    statement, err := sqlite.DB.Prepare("INSERT INTO todos (title, completed) VALUES (?, ?)")
    if err != nil {
        return nil, err
    }
    return statement.Exec(title, 0)
}

func UpdateTodoCompletion(id int, completed bool) (sql.Result, error) {
    statement, err := sqlite.DB.Prepare("UPDATE todos SET completed = ? WHERE id = ?")
    if err != nil {
        return nil, err
    }
    return statement.Exec(completed, id)
}

func DeleteTodoByID(id int) (sql.Result, error) {
    statement, err := sqlite.DB.Prepare("DELETE FROM todos WHERE id = ?")
    if err != nil {
        return nil, err
    }
    return statement.Exec(id)
}

