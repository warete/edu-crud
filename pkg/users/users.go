package users

import (
	"github.com/gin-gonic/gin"

	"github.com/warete/edu-crud/pkg/common/db"
)

func Init(r *gin.Engine, c *db.Connection) {
	RegisterRoutes(r)
}
