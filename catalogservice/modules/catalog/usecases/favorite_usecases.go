package usecases

import (
	"fmt"
	"strconv"
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"
)

type favoriteService struct {
	favRepo     entities.FavoriteRepository
	catalogRepo entities.CatalogRepository
}

func NewFavoriteService(favRepo entities.FavoriteRepository, catalogRepo entities.CatalogRepository) entities.FavoriteService {
	return favoriteService{favRepo, catalogRepo}
}

func (f favoriteService) GetFavorite(userid string) ([]entities.Favorite, error) {
	flog, err := f.favRepo.GetFavorite(userid)
	if err != nil {
		logs.Error(fmt.Sprintf("Function GetFavorite has error with user_id : %v", userid))
		return nil, err
	}
	favoriteResponse := []entities.Favorite{}
	for _, item := range flog {
		Response := entities.Favorite{
			UserID:    item.UserID,
			ID:        item.ID,
			Type:      item.Type,
			Setup:     item.Setup,
			Punchline: item.Punchline,
		}
		favoriteResponse = append(favoriteResponse, Response)
	}
	logs.Info(fmt.Sprintf("Request favorite list with account : %v", userid))
	return favoriteResponse, nil
}

func (f favoriteService) CreatedFavorite(userid string, jokeid string) (*entities.Favorite, error) {
	clog, err := f.catalogRepo.GetCatalog(jokeid)
	if err != nil {
		logs.Error(fmt.Sprintf("Function CreatedFavorite has error : %v", err))
		return nil, err
	}
	Idstr := strconv.Itoa(clog.ID)
	favorite := entities.Favorite{
		UserID:    userid,
		ID:        Idstr,
		Type:      clog.Type,
		Setup:     clog.Setup,
		Punchline: clog.Punchline,
	}
	flog, err := f.favRepo.CreateFavorite(favorite)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return flog, nil

}
