package middlewares

import (
	"my-blog-backend/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authUseCase usecase.AuthUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// トークンの取り出し
		tokenString := strings.TrimPrefix(header, "Bearer ")

		// トークンを解析してユーザーを特定
		user, err := authUseCase.GetUserFromToken(ctx, tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Controller で Get("user") が可能になる
		ctx.Set("user", user)

		println(user)

		//次の処理（Controllerなど）へ
		ctx.Next()
	}
}
