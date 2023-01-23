package service

import (
	"context"
	"errors"
	"strings"

	"github.com/tiagompalte/todo-go-sqlite/contract"
	"github.com/tiagompalte/todo-go-sqlite/entity"
)

type Service struct {
	data contract.Data
}

func NewService(data contract.Data) contract.Service {
	return &Service{
		data,
	}
}

func (s Service) List(context context.Context) (entity.Todos, error) {
	todos, err := s.data.ListAll(context)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s Service) Create(context context.Context, text string) error {
	if strings.Trim(text, "") == "" {
		return errors.New("empty text")
	}

	err := s.data.Insert(context, text)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) MarkDone(context context.Context, id uint32) error {
	todo, err := s.data.FindByID(context, id)
	if err != nil {
		return err
	}

	if todo.Done {
		return nil
	}

	todo.Done = true

	err = s.data.UpdateDone(context, todo)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Delete(context context.Context, id uint32) error {
	err := s.data.DeleteByID(context, id)
	if err != nil {
		return err
	}
	return nil
}
