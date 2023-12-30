package repositoryauthor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doublehops/dhapi-example/internal/logga"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository"
	req "github.com/doublehops/dhapi-example/internal/request"
)

type Author struct {
	Log *logga.Logga
}

func New(logger *logga.Logga) *Author {
	return &Author{
		Log: logger,
	}
}

func (a *Author) Create(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	result, err := tx.Exec(insertRecordSQL, model.UserID, model.Name, model.CreatedBy, model.UpdatedBy, model.CreatedAt, model.UpdatedAt)
	//result, err := tx.Exec(insertRecordSQL, model.UserID, model.Name, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		a.Log.Error(ctx, errMsg, nil)

		return fmt.Errorf(errMsg)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	model.ID = int32(lastInsertID)

	return nil
}

func (a *Author) Update(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	_, err := tx.Exec(updateRecordSQL, model.Name, model.UpdatedBy, model.UpdatedAt, model.ID)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		a.Log.Error(ctx, errMsg, nil)

		return fmt.Errorf(errMsg)
	}

	return nil
}

func (a *Author) Delete(ctx context.Context, tx *sql.Tx, model *model.Author) error {

	_, err := tx.Exec(deleteRecordSQL, model.UpdatedBy, model.UpdatedAt, model.DeletedAt, model.ID)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		a.Log.Error(ctx, errMsg, nil)

		return fmt.Errorf(errMsg)
	}

	return nil
}

func (a *Author) GetByID(ctx context.Context, DB *sql.DB, ID int32, model *model.Author) error {
	row := DB.QueryRow(selectByIDQuery, ID)

	err := row.Scan(&model.ID, &model.UserID, &model.Name, &model.CreatedBy, &model.UpdatedBy, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to fetch record %d", ID)
	}

	return nil
}

func (a *Author) GetAll(ctx context.Context, DB *sql.DB, p *req.Request) ([]*model.Author, error) {
	var (
		authors []*model.Author
		rows    *sql.Rows
		err     error
	)

	countQ, countParams := repository.BuildQuery(selectCollectionCountQuery, p, true)
	count, err := repository.GetRecordCount(DB, countQ, countParams)
	if err != nil {
		a.Log.Error(ctx, "GetAll()", logga.KVPs{"err": err})
	}
	p.SetRecordCount(count)

	q, params := repository.BuildQuery(selectCollectionQuery, p, false)

	a.Log.Debug(ctx, "GetAll()", logga.KVPs{"query": q})
	if len(params) == 0 {
		rows, err = DB.Query(q)
	} else {
		rows, err = DB.Query(q, params...)
	}
	if err != nil {
		a.Log.Error(ctx, "GetAll() unable to fetch rows", logga.KVPs{"err": err})
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
