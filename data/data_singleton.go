package data

import "calculator_backend/models"

var HistoryList []models.CalcHistory

func AddToHistoryList(history models.CalcHistory) {
	if len(HistoryList) < 10 {
		HistoryList = append(HistoryList, history)
	} else {
		HistoryList = append(HistoryList[:1], HistoryList[2:]...)
		HistoryList = append(HistoryList, history)
	}
}
