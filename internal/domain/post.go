package domain

type Post struct {
	ID        int
	Content   string
	MediaIDs  []int
	AuthorID  int
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
