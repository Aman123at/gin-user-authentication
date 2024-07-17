package controller

import (
	"net/http"

	helper "github.com/Aman123at/gin-userauth/dbhelper"
	tokenhelper "github.com/Aman123at/gin-userauth/helper"
	"github.com/Aman123at/gin-userauth/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WelcomeApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to User Auth"})
}

func HandleSignUp(c *gin.Context) {
	var user model.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user email is already in DB
	dbRes := helper.GetUserByMail(user.Email)

	if len(dbRes.Email) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User already registered with this email Id"})
		return
	}

	password := tokenhelper.ConvertPasswordToHash(user.Password)

	authToken := tokenhelper.GenerateToken(user.Email, user.Username, user.City)

	user.Password = password

	user.Token = authToken

	user.ID = primitive.NewObjectID()

	helper.SignUpUser(user)

	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

func HandleSignIn(c *gin.Context) {
	var user model.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check user exists or not
	userFromDB := helper.GetUserByMail(user.Email)

	// verify password
	isPasswordVerified := tokenhelper.VerifyUserPassword(userFromDB.Password, user.Password)

	if !isPasswordVerified {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Password does not match"})
		return
	}

	// generate auth token
	authToken := tokenhelper.GenerateToken(userFromDB.Email, userFromDB.Username, userFromDB.City)

	helper.UpdateTokenInDB(userFromDB.Email, authToken)

	c.JSON(http.StatusOK, gin.H{"success": true, "token": authToken})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	user := helper.GetOneUser(id)

	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

func GetUsers(c *gin.Context) {
	users := helper.GetAllUsers()

	c.JSON(http.StatusOK, gin.H{"success": true, "users": users})
}

func HandleLogoutUser(c *gin.Context) {
	id := c.Param("id")

	helper.LogoutUser(id)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User logged out successfully."})
}
