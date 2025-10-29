package services

import (
	"fmt"
	"strconv"
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"
)

type catalogService struct {
	catalogRepo   entities.CatalogRepository
	eventProducer entities.EventProducer
}

func NewCatalogService(catalogRepo entities.CatalogRepository, eventProducer entities.EventProducer) entities.CatalogService {
	return catalogService{catalogRepo, eventProducer}
}

func (c catalogService) GetCatalogs() (catalogs []entities.Catalog, err error) {
	catalogsDB, err := c.catalogRepo.GetCatalogs()
	if err != nil {
		return nil, err
	}

	// Idstr := strconv.Itoa(catalogsDB.)
	for _, n := range catalogsDB {
		catalogs = append(catalogs, entities.Catalog{
			ID:        strconv.Itoa(n.ID),
			Type:      n.Type,
			Setup:     n.Setup,
			Punchline: n.Punchline,
		})
	}
	return catalogs, nil
}

func (c catalogService) GetCatalog(userid string, jokeid string) (entities.History, error) {
	clog, err := c.catalogRepo.GetCatalog(jokeid)
	if err != nil {
		logs.Error(fmt.Sprintf("Function GetCatalog has error : %v", err))
		return entities.History{}, err
	}
	Idstr := strconv.Itoa(clog.ID)
	event := entities.History{
		UserID:    userid,
		ID:        Idstr,
		Type:      clog.Type,
		Setup:     clog.Setup,
		Punchline: clog.Punchline,
	}

	logs.Debug(fmt.Sprintf("Event : %#v", event))
	logs.Info(fmt.Sprintf("Produce with topic : %v", event.Name()))
	return event, c.eventProducer.Produce(event)
}
