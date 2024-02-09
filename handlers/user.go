package handlers

import (
	"net/http"

	"github.com/Adilfarooque/Footgo/usecase"
	"github.com/Adilfarooque/Footgo/utils/models"
	"github.com/Adilfarooque/Footgo/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			signup  body  models.UserSignUp  true	"signup"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/signup    [POST]
func UserSignUp(c *gin.Context) {
	var SignupDetail models.UserSignUp
	if err := c.ShouldBindJSON(&SignupDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(SignupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constrains not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := usecase.UsersSignUp(SignupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}

// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			login  body  models.LoginDetail  true	"login"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/userlogin     [POST]
func Userlogin(c *gin.Context) {
	var UserLoginDetails models.LoginDetail
	if err := c.ShouldBindJSON(&UserLoginDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}
	err := validator.New().Struct(UserLoginDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constrain not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := usecase.UserLogin(UserLoginDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusBadRequest, "User successfully logged in with password", user, err.Error())
	c.JSON(http.StatusBadRequest, success)
}

func ForgotPasswordSend(c *gin.Context) {
	var model models.ForgotPasswordSend
	if err := c.BindJSON(&model); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := usecase.ForgotPasswordSend(model.Phone)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "could not send OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, err.Error())
	c.JSON(http.StatusOK, success)
}
