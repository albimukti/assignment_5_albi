package api

import (
	model "finance-planner/models"
	"finance-planner/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	service *service.WalletService
}

func NewWalletHandler(service *service.WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) CreateWallet(c *gin.Context) {
	var wallet model.Wallet
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateWallet(&wallet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, wallet)
}

func (h *WalletHandler) GetWalletByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}
	wallet, err := h.service.GetWalletByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}
	c.JSON(http.StatusOK, wallet)
}

func (h *WalletHandler) ListWalletsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	wallets, err := h.service.ListWalletsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wallets)
}
