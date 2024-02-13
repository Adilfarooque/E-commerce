package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
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
func UserSignup(c *gin.Context) {
	var SignupDetails models.UserSignUp
	if err := c.ShouldBindJSON(&SignupDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	//validate the SignupDetails struct
	validate := validator.New()
	//Construct a more informative error msg based on validation errors
	if err := validate.Struct(SignupDetails); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range validationErrors {
			// Extract field name and tag from the error and construct a cleaner error message
			errorMessages = append(errorMessages, fmt.Sprintf("Field validation for %s failed on the %s tag", e.Field(), e.Tag()))
		}

		errorMessage := strings.Join(errorMessages, "\n")
		errs := response.ClientResponse(http.StatusBadRequest, errorMessage, nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := usecase.UsersSignUp(SignupDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfull signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}
