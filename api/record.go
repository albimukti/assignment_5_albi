package api

import (
	model "finance-planner/models"
	"finance-planner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordHandler struct {
	service *service.RecordService
}

func NewRecordHandler(service *service.RecordService) *RecordHandler {
	return &RecordHandler{service: service}
}

func (h *RecordHandler) CreateRecord(c *gin.Context) {
	var record model.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateRecord(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, record)
}

func (h *RecordHandler) ListRecordsByWalletID(c *gin.Context) {
	walletID := c.Param("ID")
	records, err := h.service.ListRecordsByWalletID(walletID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func (h *RecordHandler) ListRecordsByTime(c *gin.Context) {
	walletID := c.Param("ID")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	records, err := h.service.ListRecordsByTime(walletID, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}
