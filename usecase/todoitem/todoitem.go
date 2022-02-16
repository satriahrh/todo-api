package todoitemusecase

import (
	"context"

	"github.com/satriahrh/todo-api/dto"
	"github.com/satriahrh/todo-api/model"
)

type todoitem struct {
}

func New() *todoitem {
	return &todoitem{}
}

func (u *todoitem) CreateItem(context.Context, dto.CreateItemRequest) (err error) {
	panic("implement me")
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
