package repository

import (
	"context"

	"github.com/satriahrh/todo-api/model"
)

type Database interface {
	ItemInsert(ctx context.Context, item model.Item) error
	ItemGetList(ctx context.Context) ([]model.Item, error)
}
