package usecase

import (
	"context"

	"github.com/satriahrh/todo-api/dto"
	"github.com/satriahrh/todo-api/model"
)

type TodoItemUsecase interface {
	CreateItem(context.Context, dto.CreateItemRequest) (err error)
	CreateItemWithFile(context.Context, dto.CreateItemWithFileRequest) error
	GetList(ctx context.Context, isDone bool) ([]model.Item, error)
	GetItemByID(ctx context.Context, id int) (model.Item, error)
	UpdateItem(ctx context.Context, req dto.UpdateItemRequest) error
	DoneItem(ctx context.Context, id int) error
}
