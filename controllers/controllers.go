package controllers

import (
	"calculator_backend/data"
	"calculator_backend/models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetHistory(c *gin.Context) {
	c.JSON(200, models.CalcHistoryResponse{
		Data: data.HistoryList,
	})
}

func AddHistory(c *gin.Context) {
	expression := c.Query("expression")
	result := c.Query("result")

	newHistory := models.CalcHistory{
		Expression: expression,
		Result:     result,
	}

	data.AddToHistoryList(newHistory)
	log.Printf("new history: %v", data.HistoryList)

	c.JSON(200, "")
}
