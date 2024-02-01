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
	var SignupDetails models.UserSignup
	if err := c.ShouldBindJSON(&SignupDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(&SignupDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constrains not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := usecase.UsersSigUp(SignupDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}
