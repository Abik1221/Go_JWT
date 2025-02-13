package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserIdToUid(c *gin.Context, user_id string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	var err error = nil

	if userType != "USER" && user_id != uid {
		err = errors.New("you are not authorized to view this user")
	}

	return err
}
