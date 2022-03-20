package todoitemusecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/satriahrh/todo-api/dto"
	"github.com/satriahrh/todo-api/model"
	"github.com/satriahrh/todo-api/repository"
)

type todoitem struct {
	database repository.Database
	gcs      repository.Gcs
}

func New(database repository.Database, gcs repository.Gcs) *todoitem {
	return &todoitem{
		database: database,
		gcs:      gcs,
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

func (u *todoitem) CreateItemWithFile(ctx context.Context, req dto.CreateItemWithFileRequest) error {
	err := u.gcs.UploadFile(ctx, time.Now().String(), req.Document)
	return err
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
