package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/satriahrh/todo-api/model"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func New(db *gorm.DB) *database {
	return &database{
		db: db,
	}
}

func (r *database) ItemInsert(ctx context.Context, item model.Item) error {
	query := `
	INSERT INTO items
		(name, description, is_done)
	VALUES
		(@name, @description, 0)
	`
	tx := r.db.WithContext(ctx).
		Exec(query, sql.Named("name", item.Name), sql.Named("description", item.Description))

	err := tx.Error
	if err != nil {
		log.Printf("Error ItemInsert: %s\n", err.Error())
		return errors.New("something wrong with database")
	}

	return nil
}

func (r *database) ItemGetList(ctx context.Context) ([]model.Item, error) {
	panic("implement me")
}
