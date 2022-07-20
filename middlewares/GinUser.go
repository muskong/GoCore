package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoService/pkg/jwt"
)

func GinUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path

		notAuth := map[string]bool{
			"/user/login":   true,
			"/user/wechat":  true,
			"/story/list":   true,
			"/story/detail": true,
		}
		if !notAuth[path] {

			tokenName := jwt.Jwt.GetTokenName()

			token := ctx.GetHeader(tokenName)

			if token == "" {
				ctx.AbortWithStatusJSON(http.StatusOK, "token empty")
				return
			}

			err := jwt.Jwt.ValidateToken(token)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusOK, "auth error")
				return
			}
			var user jwt.Algorithm
			err = jwt.Jwt.DecodeToken(token, &user)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusOK, "auth error")
				return
			}
			ctx.Set("userId", user.Sub)
		}
		ctx.Next()
	}
}
