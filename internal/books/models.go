package books

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Year     int    `json:"year"`
	ISBN     string `json:"isbn"`
}
