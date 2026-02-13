package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/user"
)

type UserHandler struct {
	s *user.Service
}

func NewUserHandler(s *user.Service) UserHandler {
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
	err := h.s.Delete(c.Request.Context(), id)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req user.CreateParam
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	user, err := h.s.Create(c.Request.Context(), &req)
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
	var req user.UpdateParam
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	user, err := h.s.Update(c.Request.Context(), id, &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}
