package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/helpers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if bearerIsExist := strings.Contains(auth, "Bearer"); !bearerIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Unauthorized",
		})
		return
	}

	token := strings.Split(auth, " ")
	if len(token) < 2 {
		err := errors.New("Must provide Authorization header with format `Bearer {token}`")

		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	claims, err := helpers.VerifyAccessToken(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	c.Set("username", claims.Username)
	c.Set("userID", claims.UserID)

	c.Next()
}