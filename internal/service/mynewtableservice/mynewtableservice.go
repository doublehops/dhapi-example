package mynewtableservice

import (
	"context"
	"fmt"
	model "github.com/doublehops/dh-go-framework/internal/model/mynewtable"

	"github.com/doublehops/dh-go-framework/internal/app"
	"github.com/doublehops/dh-go-framework/internal/logga"
	"github.com/doublehops/dh-go-framework/internal/repository/mynewtablerepository"
	req "github.com/doublehops/dh-go-framework/internal/request"
	"github.com/doublehops/dh-go-framework/internal/service"
)

type MyNewTableService struct {
	*service.App
	myNewTableRepo *mynewtablerepository.MyNewTable
}

func New(app *service.App, mynewtableRepo *mynewtablerepository.MyNewTable) *MyNewTableService {
	return &MyNewTableService{
		App:            app,
		myNewTableRepo: mynewtableRepo,
	}
}

func (s MyNewTableService) Create(ctx context.Context, record *model.MyNewTable) (*model.MyNewTable, error) {
	ctx = context.WithValue(ctx, app.UserIDKey, 1) // todo - set this in middleware.

	if err := record.SetCreated(ctx); err != nil {
		s.Log.Error(ctx, "error in SetCreated", logga.KVPs{"error": err.Error()})
	}

	tx, _ := s.DB.BeginTx(ctx, nil)
	defer tx.Rollback() // nolint: errcheck

	err := s.myNewTableRepo.Create(ctx, tx, record)
	if err != nil {
		s.Log.Error(ctx, "unable to save new record. "+err.Error(), nil)

		return record, req.ErrCouldNotSaveRecord
	}

	err = tx.Commit()
	if err != nil {
		s.Log.Error(ctx, "unable to commit transaction"+err.Error(), nil)
	}

	a := &model.MyNewTable{}
	err = s.myNewTableRepo.GetByID(ctx, s.DB, record.ID, a)
	if err != nil {
		s.Log.Error(ctx, "unable to retrieve record. "+err.Error(), nil)
	}

	return a, nil
}

func (s MyNewTableService) Update(ctx context.Context, record *model.MyNewTable) (*model.MyNewTable, error) {
	record.SetUpdated(ctx)

	tx, _ := s.DB.BeginTx(ctx, nil)
	defer tx.Rollback() // nolint: errcheck

	err := s.myNewTableRepo.Update(ctx, tx, record)
	if err != nil {
		s.Log.Error(ctx, "unable to update record. "+err.Error(), nil)
	}

	err = tx.Commit()
	if err != nil {
		s.Log.Error(ctx, "unable to commit transaction"+err.Error(), nil)
	}

	a := &model.MyNewTable{}
	err = s.myNewTableRepo.GetByID(ctx, s.DB, record.ID, a)
	if err != nil {
		s.Log.Error(ctx, "unable to retrieve record. "+err.Error(), nil)
	}

	return a, nil
}

func (s MyNewTableService) DeleteByID(ctx context.Context, record *model.MyNewTable) error {
	tx, _ := s.DB.BeginTx(ctx, nil)
	defer tx.Rollback() // nolint: errcheck

	record.SetDeleted(ctx)

	err := s.myNewTableRepo.Delete(ctx, tx, record)
	if err != nil {
		s.Log.Error(ctx, "unable to delete record. "+err.Error(), nil)

		return fmt.Errorf("unable to delete record")
	}

	err = tx.Commit()
	if err != nil {
		s.Log.Error(ctx, "unable to commit transaction"+err.Error(), nil)
	}

	return nil
}

func (s MyNewTableService) GetByID(ctx context.Context, record *model.MyNewTable, ID int32) error {
	err := s.myNewTableRepo.GetByID(ctx, s.DB, ID, record)
	if err != nil {
		s.Log.Error(ctx, "unable to retrieve record. "+err.Error(), nil)
	}

	return nil
}

func (s MyNewTableService) GetAll(ctx context.Context, r *req.Request) ([]*model.MyNewTable, error) {
	records, err := s.myNewTableRepo.GetAll(ctx, s.DB, r)
	if err != nil {
		s.Log.Error(ctx, "unable to retrieve records. "+err.Error(), nil)
	}

	return records, nil
}
