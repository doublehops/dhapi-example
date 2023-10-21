package repositoryauthor

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/doublehops/dhapi-example/internal/logga"
	"github.com/doublehops/dhapi-example/internal/model"
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

	result, err := tx.Exec(insertRecordSQL, model.Name, model.CreatedBy, model.UpdatedBy, model.CreatedAt, model.UpdatedAt)
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

func (a *RepositoryAuthor) GetByID(ctx context.Context, DB *sql.DB, ID int32, author *model.Author) error {
	row := DB.QueryRow(selectByIDQuery, ID)

	err := row.Scan(&author.ID, &author.Name, &author.CreatedBy, &author.UpdatedBy, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to fetch record %d", ID)
	}

	return nil
}

func (a *RepositoryAuthor) GetAll(ctx context.Context, DB *sql.DB) ([]*model.Author, error) {
	var authors []*model.Author
	rows, err := DB.Query(selectAllQuery)
	if err != nil {
		return authors, fmt.Errorf("unable to fetch rows")
	}
	defer rows.Close()

	for rows.Next() {
		var record model.Author
		if err = rows.Scan(&record.ID, &record.Name, &record.CreatedBy, &record.UpdatedBy, &record.CreatedAt, &record.UpdatedAt); err != nil {
			return authors, fmt.Errorf("unable to fetch rows. %s", err)
		}

		authors = append(authors, &record)
	}

	return authors, nil
}

// populateRecord will populate model object from query.
//func (a *RepositoryAuthor) populateRecord(record *model.Author, row *sql.Rows) error {
//	err := row.Scan(&record.ID, &record.Name, &record.CreatedBy, &record.UpdatedBy, &record.CreatedAt, &record.UpdatedAt)
//
//	return err
//}
