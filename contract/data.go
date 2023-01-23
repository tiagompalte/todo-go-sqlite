package contract

import (
	"context"

	"github.com/tiagompalte/todo-go-sqlite/entity"
)

type Data interface {
	Insert(context context.Context, text string) error
	ListAll(context context.Context) (entity.Todos, error)
	UpdateDone(context context.Context, item entity.Todo) error
	FindByID(context context.Context, id uint32) (entity.Todo, error)
	DeleteByID(context context.Context, id uint32) error
}
