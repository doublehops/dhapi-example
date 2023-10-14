package repositoryauthor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doublehops/dhapi-example/internal/logga"
	"github.com/doublehops/dhapi-example/internal/model"
)

type RepositoryAuthor struct {
	DB  *sql.DB
	Log *logga.Logga
}

func New(DB *sql.DB, logger *logga.Logga) *RepositoryAuthor {
	return &RepositoryAuthor{
		DB:  DB,
		Log: logger,
	}
}

func (a RepositoryAuthor) Create(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	result, err := tx.Exec(InsertRecordSQL, model.Name, model.CreatedBy, model.UpdatedBy, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		a.Log.Error(ctx, errMsg)

		return fmt.Errorf(errMsg)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	model.ID = int32(lastInsertID)

	return nil
}

func (a RepositoryAuthor) Update(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	_, err := tx.Exec(UpdateRecordSQL, model.Name, model.UpdatedBy, model.UpdatedAt, model.ID)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		a.Log.Error(ctx, errMsg)

		return fmt.Errorf(errMsg)
	}

	return nil
}
