package main

import (
	"fmt"
	"github.com/gogaruda/seo"
)

func main() {
	slugs := []string{"hello-world", "hello-world-1"}
	slug := seo.GenerateUniqueSlug("Hello World", slugs)
	fmt.Println(slug)

	description := "HH EEHE HEHE LOLO coba saja sadalahag"
	fmt.Println(seo.Description(description, 15))

	title := "HELLO ikiasdf asn asldja sdja;sdj a;sdj"
	fmt.Println(seo.Title(title, 25))
}
