package repositories

import (
	"fmt"

	"userservice/modules/entities"
	"userservice/modules/logs"

	"gorm.io/gorm"
)

type historyRepository struct {
	db *gorm.DB
	// accRepo Account_userRepository
	// producer    services_producer.EventProducer
}

func NewHistoryRepository(db *gorm.DB) entities.HistoryRepository {
	db.Table("history").AutoMigrate(&entities.History{})
	return &historyRepository{db}
}

func (h historyRepository) GetHistory(userid string) ([]entities.History, error) {

	historys := []entities.History{}
	result := h.db.Where("user_id= ?", userid).Find(&historys)
	if result.Error != nil {
		logs.Error(fmt.Sprintf("Function GetHistory has error with user_id : %v", userid))
		return nil, result.Error

	}
	logs.Debug(fmt.Sprintf("Retrieved history with ID : %v", userid))
	return historys, nil
}

func (h historyRepository) CreateHistory(his entities.History) (*entities.History, error) {

	exec := h.db.Create(&his)
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function CreatedHistory has error : %v", exec.Error))
		return nil, exec.Error
	}
	logs.Debug(fmt.Sprintf("Created History Successfully with ID : %v", his.UserID))
	return &his, nil
}
