package scaffold

import (
	"context"
	"fmt"
)

const repositoryTemplate = "./internal/scaffold/templates/repository.tmpl"
const repositorySQLTemplate = "./internal/scaffold/templates/repositorysql.tmpl"

func (s *Scaffold) createRepository(ctx context.Context, m Model) error {

	m.ModelStructProperties = getStructProperties(m.Columns)
	m.ValidationRules = s.getValidationRules(m)
	path := fmt.Sprintf("%s/%s/%s", s.pwd, s.Config.Paths.Repository, m.LowerCase+"repository")
	repositoryFilename := fmt.Sprintf("%s/%s.go", path, m.LowerCase)
	sqlFilename := fmt.Sprintf("%s/sql.go", path)

	err := MkDir(path)
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)

		return err
	}

	// Write repository file.
	err = s.writeFile(repositoryTemplate, repositoryFilename, m)
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)

		return err
	}

	s.l.Info(ctx, "repository has been written: "+repositoryFilename, nil)

	// Write repository SQL file.
	s.ColumnSQLParams(&m)

	err = s.writeFile(repositorySQLTemplate, sqlFilename, m)
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)

		return err
	}

	s.l.Info(ctx, "repository SQL has been written: "+repositoryFilename, nil)

	return nil
}

func (s *Scaffold) ColumnSQLParams(m *Model) {

	var (
		insertCols = ""
		insertQs   = ""
		updateStmt = ""
		selectStmt = ""
	)

	for _, col := range m.Columns {
		insertCols += fmt.Sprintf("\t%s,\n", col.Original)
		insertQs += fmt.Sprintf("\t?,\n")
		updateStmt += fmt.Sprintf("\t%s=?\n", col.Original)
		selectStmt += fmt.Sprintf("\t%s,\n", col.Original)
	}

	// Remove two last chars (comma and carriage return) of each string.
	insertCols = insertCols[:len(insertCols)-2]
	insertQs = insertQs[:len(insertQs)-2]
	updateStmt = updateStmt[:len(updateStmt)-2]
	selectStmt = selectStmt[:len(selectStmt)-2]

	m.SQLCreate = insertCols
	m.SQLCreateQs = insertQs
	m.SQLUpdate = updateStmt
	m.SQLCreate = selectStmt
}
