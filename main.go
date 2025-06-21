package main

import (
	"fmt"
	"github.com/gogaruda/seo/pkg"
)

func main() {
	slugs := []string{"hello-world", "hello-world-1"}
	slug := pkg.GenerateUniqueSlug("Hello World", slugs)
	fmt.Println(slug)

	description := "HH EEHE HEHE LOLO coba saja sadalahag"
	fmt.Println(pkg.Description(description, 15))

	title := "HELLO ikiasdf asn asldja sdja;sdj a;sdj"
	fmt.Println(pkg.Title(title, 25))
}
