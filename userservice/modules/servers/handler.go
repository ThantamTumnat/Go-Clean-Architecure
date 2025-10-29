package servers

import (
	_userUsecaseConsumer "userservice/modules/consumer/handler"
	_userUsecaseHandler "userservice/modules/consumer/usecases"
	_UserHandlerProducer "userservice/modules/producer/handler"
	_UserUsecaseProducer "userservice/modules/producer/usecases"
	_controllers "userservice/modules/users/controllers"
	_usersRepository "userservice/modules/users/repositories"
	_UserUsecase "userservice/modules/users/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	// Users Group , Kafka Producer
	userGroup := v1.Group("/user")

	accRepo := _usersRepository.NewAccount_userRepository(s.Db)
	historyRepo := _usersRepository.NewHistoryRepository(s.Db)
	historyUsecase := _UserUsecase.NewHistoryService(historyRepo)

	//kafka consumer
	userEventHandler := _userUsecaseHandler.NewUserEventHandler(historyRepo)
	userConsumerHandler := _userUsecaseConsumer.NewConsumerHandler(userEventHandler)
	s.ConsumerHandler = userConsumerHandler

	//kafka produce
	eventProducer := _UserHandlerProducer.NewEventProducer(s.Producer)
	userUsercase := _UserUsecaseProducer.NewUserService(eventProducer, accRepo)

	_controllers.NewUserController(userGroup, userUsercase, historyUsecase, accRepo)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
