package scaffold

import (
	"fmt"
	"strings"
)

const modelTemplate = "./templates/model.tmpl"

func (s *Scaffold) buildStruct(ms modelStrings, columns map[string]string) string {
	var properties string
	for colName, colType := range columns {
		colName = ToPascalCase(colName)
		fmt.Printf("%s %s\n", colName, colType)
		propType := getPropertyType(colType)
		properties += fmt.Sprintf("%s %s `json:\"%s\"`\n", ms.PascalCase, propType, ms.SnakeCase)
	}

	return ""
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
