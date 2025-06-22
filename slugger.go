package seo

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Slugify membuat slug normal dari string input
func Slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	re := regexp.MustCompile(`[^\w]+`)
	s = re.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// GenerateUniqueSlug menerima input, daftar slug yang sudah ada, dan menghasilkan slug unik
func GenerateUniqueSlug(input string, existing []string) string {
	base := Slugify(input)

	// Map untuk mempercepat pengecekan eksistensi slug
	slugSet := make(map[string]bool)
	for _, s := range existing {
		slugSet[s] = true
	}

	if !slugSet[base] {
		return base
	}

	prefix := base + "-"
	suffixes := []int{}
	re := regexp.MustCompile("^" + regexp.QuoteMeta(prefix) + `(\d+)$`)

	for s := range slugSet {
		if matches := re.FindStringSubmatch(s); len(matches) == 2 {
			var n int
			fmt.Sscanf(matches[1], "%d", &n)
			suffixes = append(suffixes, n)
		}
	}

	next := 1
	if len(suffixes) > 0 {
		sort.Ints(suffixes)
		next = suffixes[len(suffixes)-1] + 1
	}

	return fmt.Sprintf("%s%d", prefix, next)
}
