package repositories

import (
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type catalogRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
	// producer    services_producer.EventProducer
}

func NewCatalogRepositoryRedis(db *gorm.DB, redisClient *redis.Client) entities.CatalogRepository {
	db.Table("catalog").AutoMigrate(&entities.CatalogDB{})
	apiData(db)
	return catalogRepositoryRedis{db, redisClient}
}

func (r catalogRepositoryRedis) GetCatalog(jokeid string) (catalogs *entities.CatalogDB, err error) {

	// key := "repository::GetCatalog" // Use the ID in the key
	key := fmt.Sprintf("repository::GetCatalog::%v", jokeid)

	// Redis Get
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &catalogs)
		if err == nil {
			logs.Debug(fmt.Sprintf("Redis : JokeId : %v", jokeid))
			return catalogs, nil
		}
	}

	//Database
	dbErr := r.db.Where("id = ?", jokeid).First(&catalogs).Error
	if dbErr != nil {
		// Check if the error is due to the record not being found
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			logs.Error(fmt.Sprintf("no record found with id : %v", jokeid))
			return nil, errors.New("no records found for the specified ID")
		}
		return catalogs, dbErr
	}

	// Redis Set
	data, err := json.Marshal(catalogs)
	if err != nil {
		logs.Error(err)
		return catalogs, err
	}

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		logs.Error(err)
		return catalogs, err
	}

	logs.Debug(fmt.Sprintf("Database : JokeId : %v", jokeid))

	return catalogs, nil
}

func (r catalogRepositoryRedis) GetCatalogId(id string) (string, error) {
	catalogs := entities.CatalogDB{}
	err := r.db.Where("id=?", id).First(&catalogs).Error
	if err != nil {
		logs.Error(err)
		return id, err
	}
	logs.Debug(fmt.Sprintf("Check Catalog Data : %v", id))
	return id, nil
}

func (r catalogRepositoryRedis) GetCatalogs() (catalogs []entities.CatalogDB, err error) {
	key := "repository::GetCatalogs"

	// Redis Get
	catalogsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(catalogsJson), &catalogs)
		if err == nil {
			query := "query"
			logs.Debug(fmt.Sprintf("Redis : %v", query))
			return catalogs, nil
		}
	}

	// Database
	err = r.db.Find(&catalogs).Error
	// Limit(30)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// Redis Set
	data, err := json.Marshal(catalogs)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}
	query := "query"
	logs.Debug(fmt.Sprintf("Database : %v", query))
	return catalogs, nil
}
