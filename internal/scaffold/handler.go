package scaffold

import (
	"context"
	"fmt"
)

const handlerTemplate = "./internal/scaffold/templates/handler.tmpl"

func (s *Scaffold) createHandler(ctx context.Context, m Model) error {

	m.ModelStructProperties = getStructProperties(m.Columns)
	m.ValidationRules = s.getValidationRules(m)
	path := fmt.Sprintf("%s/%s", s.pwd, s.Config.Paths.Handlers)
	filename := fmt.Sprintf("%s/%s.go", path, m.LowerCase)

	err := s.writeFile(handlerTemplate, filename, m)
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)

		return err
	}

	s.l.Info(context.TODO(), "handler has been written: "+filename, nil)

	return nil
}
