package entities

type Post struct {
	ID        int
	Content   string
	MediaURLs []string
	AuthorID  int
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
