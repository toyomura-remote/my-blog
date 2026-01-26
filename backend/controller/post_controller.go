package controller

import (
	dto "my-blog-backend/DTO"
	"my-blog-backend/domain"
	"my-blog-backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController interface {
	GetPosts(ctx *gin.Context)
	GetPostByDid(ctx *gin.Context)
	GetPostsByUserID(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
}

type postController struct {
	postUseCase usecase.PostUseCase
}

func NewPostController(postUseCase usecase.PostUseCase) PostController {
	return &postController{postUseCase: postUseCase}
}

func (c *postController) GetPosts(ctx *gin.Context) {
	posts, err := c.postUseCase.GetPosts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func (c *postController) GetPostByDid(ctx *gin.Context) {
	return
}

func (c *postController) GetPostsByUserID(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userID := user.(*domain.User).ID

	posts, err := c.postUseCase.GetPostsByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})

	return
}

func (c *postController) CreatePost(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID := user.(*domain.User).ID

	var input dto.CreatePostInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.postUseCase.CreatePost(ctx, input, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)

}
func (c *postController) UpdatePost(ctx *gin.Context) {
	return
}
func (c *postController) DeletePost(ctx *gin.Context) {
	return
}

//
