package handler

import (
	"crowFunding/model"
	"crowFunding/user"
	"crowFunding/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	userService user.UserService
}

func (h UserHandler) viewAllUser(gin *gin.Context) {
	user, err := h.userService.ViewAll()
	//gin.JSON(http.StatusOK, user)
	if err != nil {
		utils.HandleError(gin, http.StatusNotFound, err.Error())
		return
	}
	//gin.JSON(http.StatusOK, user)
	utils.HandleSuccess(gin, "View All User", user)
}

func (h UserHandler) addUser(gin *gin.Context) {
	var user = model.User{}
	err := gin.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("[userHandler.addUser] Error to read json bind data %v \n", err)
		utils.HandleError(gin, http.StatusUnprocessableEntity, "internal server error :" )
		return
	}

	_, err = h.userService.ViewByEmail(user.Email)
	if err == nil {
		utils.HandleError(gin, http.StatusNotFound ,"email is already registered")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Printf("[userhandler.addUser] Error to hash password %v \n", err)
		utils.HandleError(gin, http.StatusInternalServerError, "Oppss server somting wrong")
		return
	}
	user.Password = string(hash)

	_, err = h.userService.Insert(&user)

	if err != nil {
		utils.HandleError(gin, http.StatusInternalServerError, err.Error())
		return
	}

	formatter := model.UserFormatter{
		ID: user.ID,
		Name: user.Name,
		Occupation: user.Occupation,
		Email: user.Email,
		Token: "tokentokentoken",
	}


	utils.HandleSuccess(gin, "Account Has Been Register", formatter)
}

func CreateUserHandler(u *gin.Engine, userService user.UserService) {
	userHandler := UserHandler{userService}

	api := u.Group("/api/v1")
	api.GET("/userGet", userHandler.viewAllUser)
	api.POST("/user", userHandler.addUser)

}