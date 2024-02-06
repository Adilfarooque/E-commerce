package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Adilfarooque/Footgo/helper"
	"github.com/Adilfarooque/Footgo/utils/response"
	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("authorization")
		fmt.Println(tokenHeader, "this is the token header")
		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No authe header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "invalid token format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		tokenPart := splitted[1]
		tokenClaims, err := helper.ValidateToken(tokenPart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "invalid token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		c.Set("tokenClaims", tokenClaims)
		c.Next()
	}
}
