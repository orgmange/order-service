package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/user"
)

func getUintFromParam(key string, c *gin.Context) (uint, bool) {
	valRaw := c.Param(key)
	val, err := strconv.ParseUint(valRaw, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return 0, false
	}

	return uint(val), true
}

func handleAppErr(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

type UserResponse struct {
	ID    uint
	Name  string
	Email string
}

func ToUserResponse(user *user.Model) *UserResponse {
	return &UserResponse{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
}
