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
	path := fmt.Sprintf("%s/%s/%s", s.pwd, s.Config.Paths.Repository, "repository"+m.LowerCase)
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

	// Write repository SQL file.
	err = s.writeFile(repositorySQLTemplate, sqlFilename, m)
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)

		return err
	}

	s.l.Info(ctx, "repository has been written: "+repositoryFilename, nil)

	return nil
}
