package dto

type CreatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostInput struct {
	Did     string `json:"did"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetPostByDidInput struct {
	Did string `uri:"did"`
}
