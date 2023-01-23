package main

import (
	"database/sql"

	"github.com/tiagompalte/todo-go-sqlite/data"
	"github.com/tiagompalte/todo-go-sqlite/server"
	"github.com/tiagompalte/todo-go-sqlite/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repo := data.NewSqliteData(db)
	serv := service.NewService(repo)
	serve := server.NewServer(serv)

	serve.Handle()
}
