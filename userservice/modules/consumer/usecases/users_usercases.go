package usecases

import (
	"encoding/json"
	"fmt"
	"userservice/modules/entities"
	"userservice/modules/logs"
)

type userEventHandler struct {
	hisRepo entities.HistoryRepository
	// proRepo services_producer.EventProducer
}

// , proRepo services_producer.EventProducer
func NewUserEventHandler(hisRepo entities.HistoryRepository) entities.EventHandler {
	return &userEventHandler{hisRepo}
}

func (obj *userEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {

	case entities.History{}.Name():
		event := &entities.History{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(fmt.Sprintf("Consume %v has error : %v", event.Name(), err))
			return
		}
		data := entities.History{
			UserID:    event.UserID,
			ID:        event.ID,
			Type:      event.Type,
			Setup:     event.Setup,
			Punchline: event.Punchline,
		}
		create, err := obj.hisRepo.CreateHistory(data)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(fmt.Sprintf("Consume with topic : %v", event.Name()))
		logs.Debug(fmt.Sprintf("[%v] Created event : %v with Id : %v", topic, event.Name(), create.UserID))
		logs.Debug(fmt.Sprintf("Data : %v", string(eventBytes)))

		// default:
		// 	log.Printf("Unknown topic: %s", topic)
		// 	return
	}
}
