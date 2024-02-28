package scaffold

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/doublehops/dhapi-example/internal/logga"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GetFirstRune will get the first rune (char in UTF-8) from a string.
func GetFirstRune(str string) string {
	for _, c := range str {
		return string(c)
	}

	return ""
}

// ToPascalCase will convert string to pascal case.  Eg. JohnSmith.
func ToPascalCase(str string) string {
	caser := cases.Title(language.English)
	words := strings.Split(str, "_")
	for i, word := range words {
		words[i] = caser.String(word)
	}

	return strings.Join(words, "")
}

// ToKebabCase will convert string to kebab case. Eg. john-smith.
func ToKebabCase(str string) string {
	str = strings.Replace(str, "_", "-", 99)

	return str
}

// ToInitialisation will convert string its initials. Eg: JS.
func ToInitialisation(str string) string {
	initials := ""
	words := strings.Split(str, "_")
	for _, word := range words {
		initial := fmt.Sprintf("%c", word[0])
		initials += initial
	}

	return initials
}

// ToCamelCase will convert string to camel case. Eg. johnSmith
func ToCamelCase(str string) string {
	caser := cases.Title(language.English)
	words := strings.Split(str, "_")
	for i, word := range words {
		if i == 0 {
			continue
		}
		words[i] = caser.String(word)
	}

	return strings.Join(words, "")
}

// RemoveUnderscores will remove underscores from string.
func RemoveUnderscores(str string) string {
	return strings.Replace(str, "_", "", 99)
}

// CapitaliseAbbr will capitalise abbreviations of things like `id` Or `id` to `ID`.
func CapitaliseAbbr(str string) string {
	regex := "Id$"
	r := regexp.MustCompile(regex)

	return r.ReplaceAllString(str, "ID")
}

func Gofmt(filename string) error {
	cmd := exec.Command("gofmt", "-w", filename)
	return cmd.Run()
}

// writeFile will write the template and file to disk.
func (s *Scaffold) writeFile(src, dest string, tmpl Model) error {
	ctx := context.Background()

	f, err := os.Open(src)
	if err != nil {
		return errors.New("unable to open template. " + err.Error())
	}
	defer f.Close()

	source, err := io.ReadAll(f)

	f, err = os.Create(dest)
	if err != nil {
		return errors.New("unable to open destination. " + err.Error())
	}

	t, err := template.New("model").Parse(string(source))
	err = t.Execute(f, tmpl)
	if err != nil {
		return errors.New("unable to write template. " + err.Error())
	}

	if err = Gofmt(dest); err != nil {
		s.l.Warn(ctx, "unable to run fmt. "+err.Error(), logga.KVPs{"filename": dest})
	}

	return nil
}

// MkDir will recursively make the directory only if it doesn't already exist.
func MkDir(path string) error {

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	return nil
}
