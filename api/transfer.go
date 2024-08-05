package api

import (
	model "finance-planner/models"
	"finance-planner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	service *service.TransferService
}

func NewTransferHandler(service *service.TransferService) *TransferHandler {
	return &TransferHandler{service: service}
}

func (h *TransferHandler) CreateTransfer(c *gin.Context) {
	var transfer model.Transfer
	if err := c.ShouldBindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateTransfer(&transfer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transfer)
}
