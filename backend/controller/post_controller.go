package controller

import (
	"encoding/json"
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

	buf, _ := json.Marshal(posts)
	println("c", string(buf))

	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func (c *postController) GetPostByDid(ctx *gin.Context) {
	var uriInput dto.GetPostByDidInput
	if err := ctx.ShouldBindUri(&uriInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	post, err := c.postUseCase.GetPostByDid(ctx, uriInput)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": post})
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

	var bodyInput dto.CreatePostInput
	if err := ctx.ShouldBind(&bodyInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.postUseCase.CreatePost(ctx, bodyInput, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)

}

func (c *postController) UpdatePost(ctx *gin.Context) {
	_, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var bodyInput dto.UpdatePostInput
	if err := ctx.ShouldBind(&bodyInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var uriInput dto.GetPostByDidInput
	if err := ctx.ShouldBindUri(&uriInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.postUseCase.UpdatePost(ctx, bodyInput, uriInput.Did); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func (c *postController) DeletePost(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	currentUser := user.(*domain.User)

	var uriInput dto.GetPostByDidInput
	if err := ctx.ShouldBindUri(&uriInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.postUseCase.DeletePost(ctx, uriInput.Did, currentUser.ID); err != nil {
		if err.Error() == "unauthorized to delete this post" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "他人の投稿は削除できません"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "削除に失敗しました"})
		return
	}

	ctx.Status(http.StatusOK)
	return
}

//
