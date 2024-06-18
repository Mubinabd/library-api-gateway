package middleware

import (
	// "errors"

	"net/http"
	"strings"

	t "github.com/Mubinabd/library-api-gateway/token"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		url := (ctx.Request.URL.Path)

		if strings.Contains(url, "swagger") {
			ctx.Next()
			return
		} else if isValid, err := t.ValidateToken(token); !isValid && err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}

func MiddlewareAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		url := (ctx.Request.URL.Path)

		claims, err := t.ExtractClaim(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid token",
			})
			return
		}
		if strings.Contains(url, "swagger") || claims["username"] == "admin" {
			ctx.Next()
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "only admin can operate this method",
			})
			return
		}
	}
}
