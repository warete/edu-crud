package main

import (
	"github.com/gin-gonic/gin"

	"github.com/warete/edu-crud/pkg/common/db"
	"github.com/warete/edu-crud/pkg/users"
)

func main() {
	r := gin.Default()

	connection := db.Init()
	users.Init(r, connection)

	r.Run()
}
