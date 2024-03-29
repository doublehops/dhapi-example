package mynewtablerepository

import (
	"context"
	"database/sql"
	"fmt"

	model "github.com/doublehops/dh-go-framework/internal/model/mynewtable"

	"github.com/doublehops/dh-go-framework/internal/logga"
	"github.com/doublehops/dh-go-framework/internal/repository"
	req "github.com/doublehops/dh-go-framework/internal/request"
)

type MyNewTable struct {
	l *logga.Logga
}

func New(logger *logga.Logga) *MyNewTable {
	return &MyNewTable{
		l: logger,
	}
}

func (mnt *MyNewTable) Create(ctx context.Context, tx *sql.Tx, model *model.MyNewTable) error {
	result, err := tx.Exec(insertRecordSQL, model.CurrencyID, model.Name, model.CreatedAt, model.UpdatedAt, model.DeletedAt)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		mnt.l.Error(ctx, errMsg, nil)

		return fmt.Errorf(errMsg)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	model.ID = int32(lastInsertID)

	return nil
}

func (mnt *MyNewTable) Update(ctx context.Context, tx *sql.Tx, model *model.MyNewTable) error {
	_, err := tx.Exec(updateRecordSQL, model.CurrencyID, model.Name, model.CreatedAt, model.UpdatedAt, model.DeletedAt, model.ID)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		mnt.l.Error(ctx, errMsg, nil)

		return fmt.Errorf(errMsg)
	}

	return nil
}

func (mnt *MyNewTable) Delete(ctx context.Context, tx *sql.Tx, model *model.MyNewTable) error {
	_, err := tx.Exec(deleteRecordSQL, model.UpdatedAt, model.DeletedAt, model.ID)
	if err != nil {
		errMsg := fmt.Sprintf("there was an error saving record to db. %s", err)
		mnt.l.Error(ctx, errMsg, nil)

		return fmt.Errorf(errMsg)
	}

	return nil
}

func (mnt *MyNewTable) GetByID(ctx context.Context, DB *sql.DB, ID int32, record *model.MyNewTable) error {
	row := DB.QueryRow(selectByIDQuery, ID)

	err := row.Scan(&record.ID, &record.CurrencyID, &record.Name, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt)
	if err != nil {
		mnt.l.Info(ctx, "unable to fetch record", logga.KVPs{"ID": ID})

		return fmt.Errorf("unable to fetch record %d. err: %s", ID, err)
	}

	return nil
}

func (mnt *MyNewTable) GetAll(ctx context.Context, DB *sql.DB, p *req.Request) ([]*model.MyNewTable, error) {
	var (
		records []*model.MyNewTable
		rows    *sql.Rows
		err     error
	)

	countQ, countParams := repository.BuildQuery(selectCollectionCountQuery, p, true)
	count, err := repository.GetRecordCount(DB, countQ, countParams)
	if err != nil {
		mnt.l.Error(ctx, "GetAll()", logga.KVPs{"err": err})
	}
	p.SetRecordCount(count)

	q, params := repository.BuildQuery(selectCollectionQuery, p, false)

	mnt.l.Debug(ctx, "GetAll()", logga.KVPs{"query": q})
	if len(params) == 0 {
		rows, err = DB.Query(q)
	} else {
		rows, err = DB.Query(q, params...)
	}
	if err != nil {
		mnt.l.Error(ctx, "GetAll() unable to fetch rows", logga.KVPs{"err": err})
		return records, fmt.Errorf("unable to fetch rows")
	}
	defer rows.Close()
	if err != nil {
		mnt.l.Error(ctx, "GetAll() unable to fetch rows", logga.KVPs{"err": err})

		return records, fmt.Errorf("unable to fetch rows")
	}
	if rows.Err() != nil {
		mnt.l.Error(ctx, "GetAll() unable to fetch rows", logga.KVPs{"err": rows.Err()})

		return records, fmt.Errorf("unable to fetch rows")
	}

	for rows.Next() {
		var record model.MyNewTable
		if err = rows.Scan(&record.ID, &record.CurrencyID, &record.Name, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt); err != nil {
			return records, fmt.Errorf("unable to fetch rows. %s", err)
		}

		records = append(records, &record)
	}

	return records, nil
}
