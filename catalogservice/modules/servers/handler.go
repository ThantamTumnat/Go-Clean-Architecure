package servers

import (
	_catalogController "catalogservice/modules/catalog/controllers"
	_catalogRepository "catalogservice/modules/catalog/repositories"
	_catalogUsecases "catalogservice/modules/catalog/usecases"
	_CatalogHandlerConsumer "catalogservice/modules/consumer/handler"
	_CatalogUsecasesConsumer "catalogservice/modules/consumer/usecases"
	_CatalogHandlerProducer "catalogservice/modules/producer/handler"
	_CatalogUsecasesProducer "catalogservice/modules/producer/usecases"
	_redis "catalogservice/pkg/databases/redis"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	// Users Group , Kafka Producer
	catalogGroup := v1.Group("/catalog")

	accRepo := _catalogRepository.NewAccountRepository(s.Db)
	catalogRepo := _catalogRepository.NewCatalogRepositoryRedis(s.Db, _redis.InitRedis(s.Cfg))
	favRepo := _catalogRepository.NewFavoriteRepository(s.Db)

	//kafka produce / usecases
	evenProducer := _CatalogHandlerProducer.NewEventProducer(s.Producer)
	catalogUsecases := _CatalogUsecasesProducer.NewCatalogService(catalogRepo, evenProducer)
	favUsecases := _catalogUsecases.NewFavoriteService(favRepo, catalogRepo)

	//kafka consumer
	catalogHandler := _CatalogUsecasesConsumer.NewCatalogEventHandler(catalogRepo, accRepo)
	catalogConsumerHandler := _CatalogHandlerConsumer.NewConsumerHandler(catalogHandler)
	s.ConsumerHandler = catalogConsumerHandler

	_catalogController.NewCatalogController(catalogGroup, catalogUsecases, accRepo, catalogRepo)

	_catalogController.NewFavoriteHandler(catalogGroup, accRepo, catalogRepo, favUsecases)
	// _controllers.NewUserController(userGroup, userUsercase, historyUsecase, accRepo)

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
