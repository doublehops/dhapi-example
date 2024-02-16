package scaffold

import (
	"context"
	"errors"
	"github.com/doublehops/dhapi-example/internal/logga"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

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

func ToPascalCase(str string) string {
	caser := cases.Title(language.English)
	words := strings.Split(str, "_")
	for i, word := range words {
		words[i] = caser.String(word)
	}

	return strings.Join(words, "")
}

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
func MkDir(pwd, path string) error {
	dirs := strings.Split(path, "/")

	dir := pwd
	for _, d := range dirs {
		dir += "/" + d
		res, err := os.Stat(dir)
		if err != nil && !os.IsNotExist(err) {
			return errors.New("error checking directory exists. " + err.Error())
		}

		if res != nil {
			continue
		}

		err = os.Mkdir(dir, 0755)
		if err != nil {
			return errors.New("unable to make directory. " + err.Error())
		}
	}

	return nil
}
