package seo

import (
	"strings"
	"unicode/utf8"
)

func Description(input string, maxLen int) string {
	// Bersihkan input awal
	input = strings.TrimSpace(input)
	if input == "" || maxLen <= 0 {
		return ""
	}

	words := strings.Fields(input)
	var builder strings.Builder
	totalLen := 0

	for i, word := range words {
		wordLen := utf8.RuneCountInString(word)
		// Tambahkan 1 jika bukan kata pertama (untuk spasi)
		additional := wordLen
		if i > 0 {
			additional++ // spasi
		}

		// Jika penambahan kata berikutnya melebihi batas, stop
		if totalLen+additional > maxLen {
			break
		}

		if i > 0 {
			builder.WriteByte(' ')
			totalLen++
		}
		builder.WriteString(word)
		totalLen += wordLen
	}

	// Bersihkan tanda baca atau spasi menggantung
	result := strings.TrimRight(builder.String(), ",.:;!?- ")

	return result
}
