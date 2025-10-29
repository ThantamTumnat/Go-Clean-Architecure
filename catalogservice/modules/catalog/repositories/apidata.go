package repositories

import (
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func apiData(db *gorm.DB) error {

	api := "https://api.sampleapis.com/jokes/goodJokes"
	apiURL := fmt.Sprintf("%v", api)

	// Make the GET request
	response, err := http.Get(apiURL)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("error reading response body: %v", err)
	}

	// Parse JSON response
	var catalogData []entities.CatalogDB
	err = json.Unmarshal(body, &catalogData)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	// Insert jokes into the database using GORM with a check for existing records
	for _, joke := range catalogData {
		var existingCatalog entities.CatalogDB
		if err := db.Where("id = ?", joke.ID).First(&existingCatalog).Error; err != nil {
			// Record does not exist, insert it
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newJoke := entities.CatalogDB{
					ID:        joke.ID,
					Type:      joke.Type,
					Setup:     joke.Setup,
					Punchline: joke.Punchline,
				}
				if err := db.Create(&newJoke).Error; err != nil {
					logs.Error(err)
					log.Printf("Error inserting into database: %v", err)
					// Handle the error according to your application's requirements
				}
			} else {
				// Other error occurred
				logs.Error(err)
				log.Printf("Error querying database: %v", err)
				// Handle the error according to your application's requirements
			}
		} else {
			// Record already exists, skip insertion
			logs.Error(err)
			log.Printf("Record with ID %v already exists in the database\n", joke.ID)
		}
	}
	clear := "*"
	logs.Debug(fmt.Sprintf("Data inserted into the database successfully%v.", clear))
	return nil
}
