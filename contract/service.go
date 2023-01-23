package contract

import (
	"context"

	"github.com/tiagompalte/todo-go-sqlite/entity"
)

type Service interface {
	List(context context.Context) (entity.Todos, error)
	Create(context context.Context, text string) error
	MarkDone(context context.Context, id uint32) error
	Delete(context context.Context, id uint32) error
}
