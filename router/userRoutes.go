package router

import (
	"github.com/Aman123at/gin-userauth/controller"
	"github.com/Aman123at/gin-userauth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(mRoute *gin.Engine) {
	mRoute.Use(middleware.IsUserLoggedIn)

	mRoute.GET("/api/user/all", controller.GetUsers)

	mRoute.GET("/api/user/:id", controller.GetUserById)

	mRoute.GET("/api/user/logout/:id", controller.HandleLogoutUser)
}
