package router

import (
	"github.com/Aman123at/gin-userauth/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(mRoute *gin.Engine) {
	mRoute.POST("/api/signup", controller.HandleSignUp)

	mRoute.POST("/api/signin", controller.HandleSignIn)
}
