package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func RegisterRoutes(r *gin.Engine) {
	h := &handler{}

	v1 := r.Group("/api/v1")
	{
		v1.GET("users", h.getUsers)
		v1.GET("users/:id", func(ctx *gin.Context) {

		})
		v1.POST("users", func(ctx *gin.Context) {

		})
		v1.PUT("users/:id", func(ctx *gin.Context) {

		})
		v1.DELETE("users/:id", func(ctx *gin.Context) {

		})
		v1.OPTIONS("users", func(ctx *gin.Context) {

		})
	}
}

func (h handler) getUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "getUsers Called"})
}
