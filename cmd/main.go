package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/warete/edu-crud/pkg/common/db"
	"github.com/warete/edu-crud/pkg/users"
)

func main() {
	r := gin.Default()

	connection, err := db.Init("./crud.db")
	if err != nil {
		log.Fatal(err)
	}
	users.Init(r, connection)

	r.Run()
}
