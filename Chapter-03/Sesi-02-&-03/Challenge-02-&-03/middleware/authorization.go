package middleware

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"simple-rest-api-with-jwt/database"
	"simple-rest-api-with-jwt/models"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		param := c.Param("productId")
		productId, err := strconv.Atoi(param)
		if param != "" && err != nil {
			log.Println("ERROR =>", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)

		userId := int64(userData["id"].(float64))

		User := models.User{}
		Product := models.Product{}

		err = db.Debug().Select("role").First(&User, userId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		if User.Role == "user" && (c.Request.Method == "PUT" || c.Request.Method == "DELETE") {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Set("userRole", User.Role)

		if c.Request.Method == "GET" && param == "" {
			c.Next()
			return
		}

		err = db.Debug().Select("user_id").First(&Product, productId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.Role == "user" {
			if Product.UserID != userId {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error":   "unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}

		c.Next()
	}
}
