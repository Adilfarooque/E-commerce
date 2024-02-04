package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo/usecase"
	"github.com/Adilfarooque/Footgo/utils/response"
	"github.com/gin-gonic/gin"
)

func ShowAllProducts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	countStr := c.DefaultQuery("page","10")
	count , err := strconv.Atoi(countStr)
	if err != nil{
		errorRes:= response.ClientResponse(http.StatusBadRequest,"page count not in right format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	products ,err := usecase.ShowAllProducts()
}
