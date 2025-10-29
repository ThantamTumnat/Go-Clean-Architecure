package repositories

import (
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"
	"fmt"

	"gorm.io/gorm"
)

type favoriteRepository struct {
	db *gorm.DB
	// accRepo Account_userRepository
	// producer    services_producer.EventProducer
}

func NewFavoriteRepository(db *gorm.DB) entities.FavoriteRepository {
	db.Table("favorite").AutoMigrate(&entities.Favorite{})
	return favoriteRepository{db}
}

func (f favoriteRepository) GetFavorite(userid string) ([]entities.Favorite, error) {

	favorites := []entities.Favorite{}
	result := f.db.Where("user_id= ?", userid).Find(&favorites)
	if result.Error != nil {
		logs.Error(fmt.Sprintf("Function GetFavorite has error with user_id : %v", userid))
		return nil, result.Error

	}
	logs.Debug(fmt.Sprintf("Retrieved favorites list with ID : %v", userid))
	return favorites, nil
}

func (f favoriteRepository) CreateFavorite(fav entities.Favorite) (*entities.Favorite, error) {
	exec := f.db.Create(&fav)
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function CreatedFavorite has error : %v", exec.Error))
		return nil, exec.Error
	}
	logs.Debug(fmt.Sprintf("Created Favorite Successfully with ID : %v", fav.UserID))
	return &fav, nil
}
