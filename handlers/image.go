package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo/usecase"
	"github.com/Adilfarooque/Footgo/utils/response"
	"github.com/gin-gonic/gin"
)

func ShowImages(c *gin.Context) {
	product_id := c.Query("product_id")
	productID, err := strconv.Atoi(product_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	image, err := usecase.ShowImage(productID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully retrive images", image, nil)
	c.JSON(http.StatusOK, success)
}
