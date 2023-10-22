package service

import (
	"context"
	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/handlers/pagination"

	"github.com/doublehops/dhapi/resp"

	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository/repositoryauthor"
)

type AuthorService struct {
	*App
	authorRepo *repositoryauthor.RepositoryAuthor
}

func New(app *App, authorRepo *repositoryauthor.RepositoryAuthor) *AuthorService {
	return &AuthorService{
		App:        app,
		authorRepo: authorRepo,
	}
}

func (s AuthorService) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	ctx = context.WithValue(ctx, app.UserIDKey, 1) // todo - set this in middleware.

	author.SetCreated(ctx)

	tx, _ := s.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	err := s.authorRepo.Create(ctx, tx, author)
	if err != nil {
		s.Log.Error(ctx, "unable to save new record. "+err.Error())

		return author, resp.CouldNotSaveRecord
	}

	err = tx.Commit()
	if err != nil {
		s.Log.Error(ctx, "unable to commit transaction"+err.Error())
	}

	a := &model.Author{}
	err = s.authorRepo.GetByID(ctx, s.DB, author.ID, a)
	if err != nil {
		s.Log.Error(ctx, "unable to retrieve record. "+err.Error())
	}

	return a, nil
}

func (s AuthorService) Update(ctx context.Context, author *model.Author) (*model.Author, error) {
	ctx = context.WithValue(ctx, app.UserIDKey, 2) // todo - set this in middleware.

	author.SetUpdated(ctx)

	tx, _ := s.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	err := s.authorRepo.Update(ctx, tx, author)
	if err != nil {
		s.Log.Error(ctx, "unable to update record. "+err.Error())
	}

	err = tx.Commit()
	if err != nil {
		s.Log.Error(ctx, "unable to commit transaction"+err.Error())
	}

	a := &model.Author{}
	err = s.authorRepo.GetByID(ctx, s.DB, author.ID, a)
	if err != nil {
		s.Log.Error(ctx, "unable to retrieve record. "+err.Error())
	}

	return a, nil
}

func (s AuthorService) DeleteByID(ctx context.Context, author *model.Author, ID int32) error {
	tx, _ := s.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	author.SetDeleted(ctx)

	err := s.authorRepo.Delete(ctx, tx, author)
	if err != nil {
		s.Log.Error(ctx, "unable to delete record. "+err.Error())
	}

	err = tx.Commit()
	if err != nil {
		s.Log.Error(ctx, "unable to commit transaction"+err.Error())
	}

	return nil
}

func (s AuthorService) GetByID(ctx context.Context, author *model.Author, ID int32) error {
	err := s.authorRepo.GetByID(ctx, s.DB, ID, author)
	if err != nil {
		s.Log.Error(ctx, "unable to retrieve record. "+err.Error())
	}

	return nil
}

func (s AuthorService) GetAll(ctx context.Context, p *pagination.RequestPagination) ([]*model.Author, error) {
	authors, err := s.authorRepo.GetAll(ctx, s.DB, p)
	if err != nil {
		s.Log.Error(ctx, "unable to update new record. "+err.Error())
	}

	return authors, nil
}
