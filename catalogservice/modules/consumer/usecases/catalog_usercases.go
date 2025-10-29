package usecases

import (
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"
	"encoding/json"

	"fmt"
)

type catalogEventHandler struct {
	catRepo entities.CatalogRepository
	accRepo entities.AccountRepository
	// proRepo services_producer.EventProducer
}

// , proRepo services_producer.EventProducer
func NewCatalogEventHandler(catRepo entities.CatalogRepository, accRepo entities.AccountRepository) entities.EventHandler {
	return &catalogEventHandler{
		catRepo: catRepo,
		accRepo: accRepo,
		// proRepo: proRepo,
	}
}

func (obj *catalogEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {

	case entities.UserCreated{}.Name():
		event := &entities.UserCreated{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(fmt.Sprintf("Consume %v has error : %v", event.Name(), err))
			return
		}
		data := entities.Account{
			UserID:   event.UserID,
			Username: event.Username,
			Password: event.Password,
			Address:  event.Address,
		}
		create, err := obj.accRepo.CreateAccount(&data)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(fmt.Sprintf("Consume with topic : %v", event.Name()))
		logs.Debug(fmt.Sprintf("[%v] Created event : %v with user_id : %v", topic, event.Name(), create.UserID))
		logs.Debug(fmt.Sprintf("Data : %v", string(eventBytes)))

	case entities.UserUpdated{}.Name():
		event := &entities.UserUpdated{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(fmt.Sprintf("Consume %v has error : %v", event.Name(), err))
			return
		}
		data := entities.Account{
			UserID:   event.UserID,
			Username: event.Username,
			Password: event.Password,
			Address:  event.Address,
		}

		update, err := obj.accRepo.UpdateAccount(&data)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(fmt.Sprintf("Consume with topic : %v", event.Name()))
		logs.Debug(fmt.Sprintf("[%v] Updated event : %v with user_id : %v", topic, event.Name(), update.UserID))
		logs.Debug(fmt.Sprintf("Data : %v", string(eventBytes)))

	case entities.UserDeleted{}.Name():
		event := &entities.UserDeleted{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(fmt.Sprintf("Consume %v has error : %v", event.Name(), err))
			return
		}
		deleteID, err := obj.accRepo.DeleteAccount(event.UserID)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(fmt.Sprintf("Consume with topic : %v", event.Name()))
		logs.Debug(fmt.Sprintf("[%v] Deleted event : %v with user_id : %v", topic, event.Name(), deleteID))
		logs.Debug(fmt.Sprintf("Data : %v", string(eventBytes)))

	}
}
