package users

import (
	"github.com/gin-gonic/gin"

	"github.com/warete/edu-crud/pkg/common/db"
	"github.com/warete/edu-crud/pkg/users/repository"
)

type UsersInstance struct {
	uRepo       *repository.UserRepository
	uController *handler
}

func Init(r *gin.Engine, c *db.Connection) *UsersInstance {
	uRepo := repository.InitUserRepository(c)
	uController := InitController(uRepo)

	uController.RegisterRoutes(r)

	i := &UsersInstance{
		uRepo,
		uController,
	}

	return i
}
