package scaffold

import (
	"os/exec"
	"regexp"
	"strings"

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
