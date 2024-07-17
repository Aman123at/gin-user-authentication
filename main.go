package main

import (
	"fmt"

	"github.com/Aman123at/gin-userauth/router"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to auth with gin")

	// initiating a router
	route := router.Router()

	route.Use(gin.Logger())

	route.Run(":4000")
}
