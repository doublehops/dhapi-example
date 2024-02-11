package scaffold

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/doublehops/go-common/str"
)

const modelTemplate = "./internal/scaffold/templates/model.tmpl"

func (s *Scaffold) createModel(tmpl Template, columns map[string]string) error {

	tmpl.ModelStructProperties = getPropertyTypes(columns)
	tmpl.ValidationRules = s.getValidationRules(tmpl, columns)
	filename := fmt.Sprintf("%s.go", tmpl.LowerCase)

	err := s.writeFile(filename, tmpl)
	if err != nil {
		return err
	}

	return nil
}

func getPropertyTypes(columns map[string]string) string {
	var properties string

	ignoreColumns := []string{"created_at", "updated_at", "deleted_at"}

	for colName, colType := range columns {
		if str.SliceContains(colName, ignoreColumns) {
			continue
		}

		colName = ToPascalCase(colName)
		propType := getPropertyType(colType)
		jsonVal := ToCamelCase(colName)

		properties += fmt.Sprintf("%s %s `json:\"%s\"`\n", colName, propType, jsonVal)
	}

	return properties
}

func (s *Scaffold) getValidationRules(tmpl Template, columns map[string]string) string {
	var rules string

	ignoreColumns := []string{"created_at", "updated_at", "deleted_at"}

	for colName, _ := range columns {
		if str.SliceContains(colName, ignoreColumns) {
			continue
		}

		colName = ToPascalCase(colName)

		rules += fmt.Sprintf("{\"%s\", %s.%s, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, \"\")}},\n", colName, tmpl.FirstInitial, colName)
	}

	return rules
}

// getPropertyType will check which column type the property is and return a corresponding
// Go Type to use in the model's struct.
func getPropertyType(propType string) string {
	if strings.Contains(propType, "int") {
		return "int"
	}
	if strings.Contains(propType, "char") {
		return "string"
	}
	if strings.Contains(propType, "datetime") {
		return "*time.Time"
	}

	return "string"
}

func (s *Scaffold) writeFile(filename string, tmpl Template) error {
	src := fmt.Sprintf(modelTemplate)
	f, err := os.Open(src)
	if err != nil {
		return errors.New("unable to open template. " + err.Error())
	}

	source, err := io.ReadAll(f)

	dest := fmt.Sprintf("%s/%s/%s", s.pwd, s.Config.Paths.Model, filename)
	f, err = os.Create(dest)
	if err != nil {
		return errors.New("unable to open destination. " + err.Error())
	}

	t, err := template.New("model").Parse(string(source))
	err = t.Execute(f, tmpl)
	if err != nil {
		return errors.New("unable to write template. " + err.Error())
	}

	return nil
}
