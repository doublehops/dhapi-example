package scaffold

import (
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
