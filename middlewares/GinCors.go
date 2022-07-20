package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinCORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("Origin") != "" {
			switch ctx.Request.Method {
			case http.MethodOptions:
				ctx.Header("Access-Control-Allow-Origin", "*")  // * can be replaced with any domain name which you want support cross-domain
				ctx.Header("Access-Control-Allow-Methods", "*") //enable cross domain
				ctx.Header("Access-Control-Allow-Headers", "*") //Header
			default:
				ctx.Header("Access-Control-Allow-Origin", "*") // * can be replaced with any domain name which you want support cross-domain
			}
		}
		ctx.Next()
	}
}
