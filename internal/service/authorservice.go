package service

import (
	"context"
	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository/repositoryauthor"
)

type AuthorService struct {
	app        *app.App
	authorRepo *repositoryauthor.RepositoryAuthor
}

func New(app *app.App, authorRepo *repositoryauthor.RepositoryAuthor) *AuthorService {
	return &AuthorService{
		app:        app,
		authorRepo: authorRepo,
	}
}

func (s AuthorService) Create(ctx context.Context, author *model.Author) (model.Author, error) {
	ctx = context.WithValue(ctx, app.UserIDKey, 1) // todo - set this in middleware.

	author.SetCreated(ctx)

	tx, _ := s.app.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	err := s.authorRepo.Create(ctx, tx, author)
	if err != nil {
		s.app.Log.Error(ctx, "unable to save new record. "+err.Error())
	}

	err = tx.Commit()
	if err != nil {
		s.app.Log.Error(ctx, "unable to commit transaction"+err.Error())
	}

	return *author, nil
}
