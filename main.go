package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/API-multi-level-user-golang/database"
	"github.com/haviz000/API-multi-level-user-golang/route"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	database.ConnectDB()
	db := database.GetDB()

	route.Route(router, db)
	router.Run(PORT)
}
