package dto

type CreatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
