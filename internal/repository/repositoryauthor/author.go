package repositoryauthor

import (
	"log/slog"
	"reflect"

	"github.com/doublehops/dhapi-example/internal/model"
)

type RepositoryAuthor struct {
	DB  *sql.DB
	log *slog.Logger
}

func New(DB *sql.DB, logger *slog.Logger) *RepositoryAuthor {
	return &RepositoryAuthor{
		DB:  DB,
		log: logger,
	}
}

func (a RepositoryAuthor) Create(author *model.Author) (*model.Author, error) {
	a.log.Info("Adding author: %s; with interface type: %v", author.Name, reflect.TypeOf(m.db))

	result, err := a.DB.Exec(InsertRecordSQL, record.Name, record.Symbol)
	if err != nil {
		a.log.Error("There was an error saving record to db. %s", err)

		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
