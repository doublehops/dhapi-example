package scaffold

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"text/template"
)

const routeTemplate = "./internal/scaffold/templates/routes.tmpl"

// printRoutes will print the routes for the user to manually add to the routes table.
func (s *Scaffold) printRoutes(ctx context.Context, m Model) error {

	m.ModelStructProperties = getStructProperties(m.Columns)

	f, err := os.Open(routeTemplate)
	if err != nil {
		e := "unable to open template. " + err.Error()
		s.l.Error(ctx, e, nil)

		return errors.New(e)
	}
	defer f.Close()

	source, err := io.ReadAll(f)

	buf := bytes.Buffer{}

	t, err := template.New("model").Parse(string(source))
	err = t.Execute(&buf, m)
	if err != nil {
		e := "unable to write new file. " + err.Error()
		s.l.Error(ctx, e, nil)

		return errors.New(e)
	}

	fmt.Print(buf.String())

	return nil
}
