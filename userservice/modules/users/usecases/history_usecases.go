package usecases

import (
	"fmt"
	"userservice/modules/entities"
	"userservice/modules/logs"
)

type historyService struct {
	hisRepo entities.HistoryRepository
}

func NewHistoryService(hisRepo entities.HistoryRepository) entities.HistoryService {
	return historyService{hisRepo}
}

func (h historyService) GetHistory(userid string) ([]entities.History, error) {
	hlog, err := h.hisRepo.GetHistory(userid)
	if err != nil {
		logs.Error(fmt.Sprintf("Function GetHistory has error with user_id : %v", userid))
		return nil, err
	}
	historyResponse := []entities.History{}
	for _, item := range hlog {
		Response := entities.History{
			UserID:    item.UserID,
			ID:        item.ID,
			Type:      item.Type,
			Setup:     item.Setup,
			Punchline: item.Punchline,
		}
		historyResponse = append(historyResponse, Response)
	}
	logs.Info(fmt.Sprintf("Request account history with user_id : %v", userid))
	return historyResponse, nil
}
