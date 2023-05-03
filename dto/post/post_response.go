package post

type PostResponse struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Photo       string       `json:"photo"`
	UserID      int          `json:"user_id"`
}