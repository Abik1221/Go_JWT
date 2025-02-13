package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HashPassward() {

}

func VerifyPassword() {

}

func Signup() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		user_id := c.Param("id")

		err := helper.MatchUserIdToUid(c, user_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
