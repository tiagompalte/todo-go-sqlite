package data

import (
	"context"
	"database/sql"

	"github.com/tiagompalte/todo-go-sqlite/contract"
	"github.com/tiagompalte/todo-go-sqlite/entity"
)

type SqliteData struct {
	db *sql.DB
}

func NewSqliteData(db *sql.DB) contract.Data {
	sts := "CREATE TABLE IF NOT EXISTS todo(id INTEGER PRIMARY KEY, text TEXT NOT NULL, done BOOL DEFAULT FALSE);"
	_, err := db.Exec(sts)
	if err != nil {
		panic(err)
	}
	return SqliteData{db}
}

func (s SqliteData) Insert(context context.Context, text string) error {
	_, err := s.db.ExecContext(context, "INSERT INTO todo(text) VALUES(?);", text)
	if err != nil {
		return err
	}

	return nil
}

func (s SqliteData) ListAll(context context.Context) (entity.Todos, error) {
	rows, err := s.db.QueryContext(context, "SELECT id, text, done FROM todo")
	if err != nil {
		return nil, err
	}

	ret := make(entity.Todos, 0)
	for rows.Next() {
		var todo entity.Todo

		err = rows.Scan(&todo.ID, &todo.Text, &todo.Done)
		if err != nil {
			return nil, err
		}

		ret = append(ret, todo)
	}

	return ret, nil
}

func (s SqliteData) UpdateDone(context context.Context, item entity.Todo) error {
	_, err := s.db.ExecContext(context, "UPDATE todo SET done = ? WHERE id = ?;", item.Done, item.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s SqliteData) FindByID(context context.Context, id uint32) (entity.Todo, error) {
	row := s.db.QueryRowContext(context, "SELECT id, text, done FROM todo WHERE id = ?", id)

	var todo entity.Todo
	err := row.Scan(&todo.ID, &todo.Text, &todo.Done)
	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}

func (s SqliteData) DeleteByID(context context.Context, id uint32) error {
	_, err := s.db.ExecContext(context, "DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
