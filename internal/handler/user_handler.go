package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/dto"
	"github.com/orgmange/order-service/internal/service"
)

type UserHandler struct {
	s service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return UserHandler{
		s: s,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, ok := getIntFromParam("id", c)
	if !ok {
		return
	}
	user, err := h.s.GetUser(id)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, ok := getIntFromParam("id", c)
	if !ok {
		return
	}
	err := h.s.DeleteUser(id)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	user, err := h.s.CreateUser(&req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, ok := getIntFromParam("id", c)
	if !ok {
		return
	}
	var req dto.UpdateUserRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	user, err := h.s.UpdateUser(id, &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

func getIntFromParam(key string, c *gin.Context) (int, bool) {
	valRaw := c.Param(key)
	val, err := strconv.Atoi(valRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return 0, false
	}

	return val, true
}

func handleAppErr(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
