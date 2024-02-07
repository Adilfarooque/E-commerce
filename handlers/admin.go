package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo/usecase"
	"github.com/Adilfarooque/Footgo/utils/models"
	"github.com/Adilfarooque/Footgo/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary		Admin Login
// @Description	Login handler for jerseyhub admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AdminLogin	true	"Admin login details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/adminlogin [POST]
func LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := usecase.LoginHandler(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Can't autheticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin authenticate successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Admin Dashboard
// @Description	Retrieve admin dashboard
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/dashboard [GET]
func DashBoard(c *gin.Context) {
	adminDashboard, err := usecase.DashBoard()
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin Dashboard displayed", adminDashboard, err.Error())
	c.JSON(http.StatusOK, success)
}

// @Summary		Get Users
// @Description	Retrieve users with pagination
// @Tags			Admin User Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param page query string false "Page number"
// @Param count query string false "Page size"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/users   [GET]\
func GetUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		erroRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, erroRes)
		return
	}
	countStr := c.DefaultQuery("count", "10")
	pageSize, err := strconv.Atoi(countStr)
	if err != nil {
		erroRes := response.ClientResponse(http.StatusBadRequest, "user couint in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, erroRes)
		return
	}
	users, err := usecase.ShowAllUsers(page, pageSize)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "couldn't retrieve users", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Retrived all users", users, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Block User
// @Description	using this handler admins can block an user
// @Tags			Admin User Management
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/users/block   [PUT]
func BlockUser(c *gin.Context) {
	id := c.Query("id")
	err := usecase.BlockedUser(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "user couldn't be blocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, err.Error())
	c.JSON(http.StatusOK, success)
}

// @Summary		UnBlock an existing user
// @Description	UnBlock user
// @Tags			Admin User Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/users/unblock    [PUT]

func UnBlockUser(c *gin.Context) {
	id := c.Query("id")
	err := usecase.UnBlockedUser(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "user could not be unblocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, err.Error())
	c.JSON(http.StatusOK, success)
}

// @Summary Filtered Sales Report
// @Description Get Filtered sales report by week, month and year
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param period query string true "sales report"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/sales-report    [GET]
/*
func FilteredSalesReport(c *gin.Context) {
	timePeriod := c.Query("period")
	salesReport, err := usecase.FilteredSalesReport()
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "report couldn't retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	success := response.ClientResponse(http.StatusOK, "report retrieved successfully", salesReport, nil)
	c.JSON(http.StatusOK, success)
}
*/