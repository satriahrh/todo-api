package todoitemusecase

import (
	"context"
	"errors"
	"log"

	"github.com/satriahrh/todo-api/dto"
	"github.com/satriahrh/todo-api/model"
	"github.com/satriahrh/todo-api/repository"
)

type todoitem struct {
	database repository.Database
}

func New(database repository.Database) *todoitem {
	return &todoitem{
		database: database,
	}
}

func (u *todoitem) CreateItem(ctx context.Context, req dto.CreateItemRequest) (err error) {
	if req.Name == "" {
		return errors.New("name tidak boleh kosong")
	}
	if len(req.Description) > 255 {
		return errors.New("tidak boleh terlalu panjang")
	}

	item := model.Item{
		Name:        req.Name,
		Description: req.Description,
	}

	err = u.database.ItemInsert(ctx, item)
	if err != nil {
		log.Printf("Error in CreateItem: %s\n", err.Error())
		return err
	}

	return nil
}

func (u *todoitem) GetList(ctx context.Context, isDone bool) ([]model.Item, error) {
	panic("implement me")
}

func (u *todoitem) GetItemByID(ctx context.Context, id int) (model.Item, error) {
	panic("implement me")
}

func (u *todoitem) UpdateItem(ctx context.Context, req dto.UpdateItemRequest) error {
	panic("implement me")
}

func (u *todoitem) DoneItem(ctx context.Context, id int) error {
	panic("implement me")
}
