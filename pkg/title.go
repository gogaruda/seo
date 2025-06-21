package pkg

import (
	"strings"
	"unicode"
)

func Title(s string, maxLen int) string {
	s = strings.TrimSpace(s)
	if s == "" || maxLen <= 0 {
		return ""
	}

	words := strings.Fields(s)
	var result []string
	totalLen := 0

	for i, word := range words {
		runes := []rune(word)

		if len(runes) == 0 {
			continue
		}

		// Huruf pertama kapital, sisanya kecil
		runes[0] = unicode.ToUpper(runes[0])
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}

		capitalized := string(runes)

		additional := len([]rune(capitalized))
		if i > 0 {
			additional++ // spasi
		}

		if totalLen+additional > maxLen {
			break
		}

		if i > 0 {
			totalLen++ // spasi
		}
		result = append(result, capitalized)
		totalLen += len([]rune(capitalized))
	}

	title := strings.Join(result, " ")
	title = strings.TrimRight(title, ",.:;!?- ")

	return title
}
