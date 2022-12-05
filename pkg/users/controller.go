package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/warete/edu-crud/pkg/users/repository"
)

type handler struct {
	uRepo *repository.UserRepository
}

func InitController(uRepo *repository.UserRepository) *handler {
	return &handler{
		uRepo,
	}
}

func (h handler) RegisterRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		v1.GET("users", h.getUsers)
		v1.GET("users/:id", h.getUserById)
		v1.POST("users", h.addUser)
		v1.PUT("users/:id", h.updateUser)
		v1.DELETE("users/:id", h.deleteUser)
	}
}

func (h handler) getUsers(ctx *gin.Context) {
	users, err := h.uRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	if users == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func (h handler) getUserById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong id param"})
	}

	user, err := h.uRepo.GetById(id)

	if err != nil {
		log.Fatal(err)
	}
	if user.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func (h handler) addUser(ctx *gin.Context) {

	var json repository.User

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := h.uRepo.Add(json)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"data": success})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func (h handler) updateUser(ctx *gin.Context) {

	var json repository.User

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong id param"})
	}

	err = h.uRepo.Update(id, json)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func (h handler) deleteUser(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong id param"})
	}

	err = h.uRepo.Delete(id)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
