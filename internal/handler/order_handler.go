package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/dto"
	"github.com/orgmange/order-service/internal/service"
)

type OrderHandler struct {
	s service.OrderService
}

func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{
		s: s,
	}
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, ok := getUintFromParam("id", c)
	if !ok {
		return
	}

	order, err := h.s.GetOrder(c.Request.Context(), id)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	order, err := h.s.CreateOrder(c.Request.Context(), &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var req dto.UpdateOrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = h.s.UpdateOrder(c.Request.Context(), &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, ok := getUintFromParam("id", c)
	if !ok {
		return
	}

	err := h.s.DeleteOrder(c.Request.Context(), id)
	if err != nil {
		handleAppErr(err, c)
		return
	}
}
