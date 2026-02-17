package controllers

import (
	"net/http"
	"strconv"

	"github.com/busychirag/assignments/tree/main/Assignment2/services"
	"github.com/gin-gonic/gin"
)

func GetTransactionsByAccountID(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	transactions, err := services.GetTransactionsByAccountID(uint64(accountID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
