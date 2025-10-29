package usecases

import (
	"fmt"
	"userservice/modules/entities"
	"userservice/modules/logs"

	"github.com/google/uuid"
)

type userService struct {
	eventProducer entities.EventProducer
	accRepo       entities.Account_userRepository
}

func NewUserService(eventProducer entities.EventProducer, accRepo entities.Account_userRepository) entities.UserService {
	return &userService{eventProducer: eventProducer, accRepo: accRepo}
}

func (obj userService) UserCreated(command *entities.UserCreated) (entities.UserCreated, error) {

	data := &entities.Account_user{
		UserID:   uuid.NewString(),
		Username: command.Username,
		Password: command.Password,
		Address:  command.Address,
	}
	created, err := obj.accRepo.CreateAccount(data)
	if err != nil {
		logs.Error(err)
		return entities.UserCreated{}, err
	}

	event := entities.UserCreated{
		UserID:   created.UserID,
		Username: created.Username,
		Password: created.Password,
		Address:  created.Address,
	}
	logs.Debug(fmt.Sprintf("Event : %#v", event))
	logs.Info(fmt.Sprintf("Produce with topic : %v", event.Name()))
	return event, obj.eventProducer.Produce(event)

}

func (obj userService) UserUpdated(command *entities.UserUpdated) (entities.UserUpdated, error) {
	data := &entities.Account_user{
		UserID:   command.UserID,
		Username: command.Username,
		Password: command.Password,
		Address:  command.Address,
	}
	updated, err := obj.accRepo.UpdateAccount(data)
	if err != nil {
		logs.Error(err)
		return entities.UserUpdated{}, err
	}

	event := entities.UserUpdated{
		UserID:   updated.UserID,
		Username: updated.Username,
		Password: updated.Password,
		Address:  updated.Address,
	}
	logs.Debug(fmt.Sprintf("Event : %#v", event))
	logs.Info(fmt.Sprintf("Produce with topic : %v", event.Name()))
	return event, obj.eventProducer.Produce(event)
}

func (obj userService) UserDeleted(command *entities.UserDeleted) (entities.UserDeleted, error) {
	// if command.UserID == "" {
	// 	return events.UserDeleted{}, errors.New("user_id could not empty")
	// }
	data := entities.Account_user{
		UserID: command.UserID,
	}
	deleted, err := obj.accRepo.DeleteAccount(data.UserID)
	if err != nil {
		logs.Error(err)
		return entities.UserDeleted{}, err
	}

	event := entities.UserDeleted{
		UserID: deleted,
	}
	logs.Debug(fmt.Sprintf("Event : %#v", event))
	logs.Info(fmt.Sprintf("Produce with topic : %v", event.Name()))
	return event, obj.eventProducer.Produce(event)
}
