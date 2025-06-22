# Simple SEO

## Install
```
go get github.com/gogaruda/seo@v1.0.1
```

## Penggunaan
### Slugger
```go
func handleCreatePost(input string, db *sql.DB) (string, error) {
	base := slugger.Slugify(input)

	rows, err := db.Query(`SELECT slug FROM posts WHERE slug LIKE ?`, base+"%")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var existing []string
	for rows.Next() {
		var slug string
		rows.Scan(&slug)
		existing = append(existing, slug)
	}

	uniqueSlug := slugger.GenerateUniqueSlug(input, existing)

	// simpan ke DB
	_, err = db.Exec(`INSERT INTO posts (title, slug) VALUES (?, ?)`, input, uniqueSlug)
	if err != nil {
		return "", err
	}

	return uniqueSlug, nil
}
```