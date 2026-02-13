package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/order"
)

type OrderResponse struct {
	ID        uint
	CreatorID uint
	Status    string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func ToOrderResponse(order *order.Model) *OrderResponse {
	return &OrderResponse{
		ID:        order.GetID(),
		CreatorID: order.GetCreatorID(),
		Status:    string(order.GetStatus()),
		UpdatedAt: order.GetUpdatedAt(),
		CreatedAt: order.GetCreatedAt(),
	}
}

type OrderHandler struct {
	s *order.Service
}

func NewOrderHandler(s *order.Service) *OrderHandler {
	return &OrderHandler{
		s: s,
	}
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, ok := getUintFromParam("id", c)
	if !ok {
		return
	}

	order, err := h.s.Get(c.Request.Context(), id)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req order.CreateParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	order, err := h.s.Create(c.Request.Context(), &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id, ok := getUintFromParam("id", c)
	if !ok {
		return
	}
	var req order.UpdateParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	order, err := h.s.Update(c.Request.Context(), id, &req)
	if err != nil {
		handleAppErr(err, c)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, ok := getUintFromParam("id", c)
	if !ok {
		return
	}

	err := h.s.Delete(c.Request.Context(), id)
	if err != nil {
		handleAppErr(err, c)
		return
	}
}
