package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"lokucrazy/Go-DnD-Info/models"
)

type racesService struct {
	DB *sql.DB
}

func (r *racesService) Get(name string) ([]byte, error) {
	var races []models.Races
	query := "SELECT * FROM races "
	if name != "" {
		query += "WHERE races.name = ?"
	}

	result, err := r.DB.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer closeRows(result)

	for result.Next() {
		race := models.Races{}
		err = result.Scan(
			&race.ID,
			&race.Name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		races = append(races, race)
	}

	return json.Marshal(races)
}

func CreateRacesService(db *sql.DB) *racesService {
	return &racesService{DB: db}
}
