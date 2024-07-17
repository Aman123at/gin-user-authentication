package router

import (
	"github.com/Aman123at/gin-userauth/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()

	// welcome route
	r.GET("/", controller.WelcomeApi)

	// login , signup routes (does not require authentication token)
	AuthRoutes(r)

	// user routes (require authentication token)
	UserRoutes(r)

	return r
}
