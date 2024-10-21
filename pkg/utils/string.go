package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Функция сложения строк (оптимизация стандартной конкатенации)
func ConcatStrings(s ...string) (string, error) {
	var builder strings.Builder

	for _, str := range s {
		_, err := builder.WriteString(str)
		if err != nil {
			return "", fmt.Errorf("error concat string %s", str)
		}
	}

	str := builder.String()

	return str, nil
}

// Добавляем экранирование для спец символов MarkdownV2, чтобы Telegram смог распарсить текст
func MarkDownFilter(text string) string {
	markdownSpecialChars := regexp.MustCompile(`([_\*\[\]\(\)~` + "`" + `>#+\-=|{}.!])`)

	replaceFn := func(char string) string {
		return `\` + char
	}

	return markdownSpecialChars.ReplaceAllStringFunc(text, replaceFn)
}
