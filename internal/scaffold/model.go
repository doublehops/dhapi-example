package scaffold

import (
	"context"
	"fmt"
	"github.com/doublehops/go-common/str"
	"strings"
)

const modelTemplate = "./internal/scaffold/templates/model.tmpl"

func (s *Scaffold) createModel(ctx context.Context, m Model) error {

	m.ModelStructProperties = getStructProperties(m.Columns)
	m.ValidationRules = s.getValidationRules(m)
	path := fmt.Sprintf("%s/%s/%s", s.pwd, s.Config.Paths.Model, m.LowerCase)
	filename := fmt.Sprintf("%s/%s.go", path, m.LowerCase)

	err := MkDir(path)
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)

		return err
	}

	err = s.writeFile(modelTemplate, filename, m)
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)

		return err
	}

	s.l.Info(ctx, "model has been written: "+filename, nil)

	return nil
}

func getStructProperties(columns []column) string {

	ignoreColumns := []string{"created_at", "updated_at", "deleted_at"}

	var properties string

	for _, col := range columns {
		if str.SliceContains(col.Original, ignoreColumns) {
			continue
		}

		properties += fmt.Sprintf("%s %s `json:\"%s\"`\n", col.CapitalisedAbbr, col.Type, col.CamelCase)
	}

	return properties
}

func (s *Scaffold) getValidationRules(m Model) string {
	var rules string

	ignoreColumns := []string{"created_at", "updated_at", "deleted_at"}

	for _, col := range m.Columns {
		if str.SliceContains(col.Original, ignoreColumns) {
			continue
		}

		rules += fmt.Sprintf("{\"%s\", %s.%s, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, \"\")}},\n", col.CamelCase, m.FirstInitial, col.CapitalisedAbbr)
	}

	return rules
}

// getPropertyType will check which column type the property is and return a corresponding
// Go Type to use in the model's struct.
func getPropertyType(propType string) columnType {
	if strings.Contains(propType, "int") {
		return typeInt
	}
	if strings.Contains(propType, "char") {
		return typeString
	}
	if strings.Contains(propType, "datetime") {
		return typeDatetime
	}

	return typeString
}

//func (s *Scaffold) writeModelFile(filename string, tmpl Model) error {
//	src := fmt.Sprintf(modelTemplate)
//	f, err := os.Open(src)
//	if err != nil {
//		return errors.New("unable to open template. " + err.Error())
//	}
//
//	source, err := io.ReadAll(f)
//
//	dest := fmt.Sprintf("%s/%s/%s", s.pwd, s.Config.Paths.Model, filename)
//	f, err = os.Create(dest)
//	if err != nil {
//		return errors.New("unable to open destination. " + err.Error())
//	}
//
//	t, err := template.New("model").Parse(string(source))
//	err = t.Execute(f, tmpl)
//	if err != nil {
//		return errors.New("unable to write template. " + err.Error())
//	}
//
//	if err = Gofmt(dest); err != nil {
//		return errors.New("unable to run gofmt. " + err.Error())
//	}
//
//	return nil
//}
