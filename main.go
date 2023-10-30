package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
    // Inicialize o banco de dados SQLite
    var err error
    db, err = sql.Open("sqlite3", "./app.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Crie a tabela de exemplo
    createTable()

    // Configure as rotas da aplicação
    r := mux.NewRouter()
    r.HandleFunc("/items", getItems).Methods("GET")
    r.HandleFunc("/items", addItem).Methods("POST")

    // Inicie o servidor HTTP
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTable() {
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );`
    
    _, err := db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}

func getItems(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM items")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var items []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            log.Fatal(err)
        }
        items = append(items, name)
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"items": %q}`, items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    _, err := db.Exec("INSERT INTO items (name) VALUES (?)", name)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"message": "Item added successfully"}`)
}
