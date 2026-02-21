package controller

import (
	dto "my-blog-backend/DTO"
	"my-blog-backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	usecase usecase.AuthUseCase
}

func NewAuthController(usecase usecase.AuthUseCase) AuthController {
	return &authController{usecase: usecase}
}

func (c *authController) SignUp(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.usecase.SignUp(ctx, input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user":  response.User,
		"token": response.Token,
	})
}

func (c *authController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.usecase.Login(ctx, input.Email, input.Password)
	if err != nil {
		if err.Error() == "User not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user":  response.User,
		"token": response.Token,
	})
}
