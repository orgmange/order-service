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
	id, ok := getUintFromParam("id", c)
	if !ok {
		return
	}
	user, err := h.s.GetUser(c.Request.Context(), id)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, ok := getUintFromParam("id", c)
	if !ok {
		return
	}
	err := h.s.DeleteUser(c.Request.Context(), id)
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
	user, err := h.s.CreateUser(c.Request.Context(), &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, ok := getUintFromParam("id", c)
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
	user, err := h.s.UpdateUser(c.Request.Context(), id, &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

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
