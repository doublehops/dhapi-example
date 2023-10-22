package repositoryauthor

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/doublehops/dhapi-example/internal/handlers/pagination"
	"github.com/doublehops/dhapi-example/internal/logga"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository"
)

type RepositoryAuthor struct {
	Log *logga.Logga
}

func New(logger *logga.Logga) *RepositoryAuthor {
	return &RepositoryAuthor{
		Log: logger,
	}
}

func (a *RepositoryAuthor) Create(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	result, err := tx.Exec(insertRecordSQL, model.UserID, model.Name, model.CreatedBy, model.UpdatedBy, model.CreatedAt, model.UpdatedAt)
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

func (a *RepositoryAuthor) Update(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	_, err := tx.Exec(updateRecordSQL, model.Name, model.UpdatedBy, model.UpdatedAt, model.ID)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		a.Log.Error(ctx, errMsg)

		return fmt.Errorf(errMsg)
	}

	return nil
}

func (a *RepositoryAuthor) Delete(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	_, err := tx.Exec(deleteRecordSQL, model.UpdatedBy, model.UpdatedAt, model.DeletedAt, model.ID)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		a.Log.Error(ctx, errMsg)

		return fmt.Errorf(errMsg)
	}

	return nil
}

func (a *RepositoryAuthor) GetByID(ctx context.Context, DB *sql.DB, ID int32, model *model.Author) error {
	row := DB.QueryRow(selectByIDQuery, ID)

	err := row.Scan(&model.ID, &model.UserID, &model.Name, &model.CreatedBy, &model.UpdatedBy, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to fetch record %d", ID)
	}

	return nil
}

func (a *RepositoryAuthor) GetAll(ctx context.Context, DB *sql.DB, p *pagination.RequestPagination) ([]*model.Author, error) {
	var authors []*model.Author
	q := repository.SubstitutePaginationVars(selectAllQuery, p)
	a.Log.Info(ctx, q)
	rows, err := DB.Query(q)
	if err != nil {
		return authors, fmt.Errorf("unable to fetch rows")
	}
	defer rows.Close()

	for rows.Next() {
		var record model.Author
		if err = rows.Scan(&record.ID, &record.UserID, &record.Name, &record.CreatedBy, &record.UpdatedBy, &record.CreatedAt, &record.UpdatedAt); err != nil {
			return authors, fmt.Errorf("unable to fetch rows. %s", err)
		}

		authors = append(authors, &record)
	}

	return authors, nil
}
