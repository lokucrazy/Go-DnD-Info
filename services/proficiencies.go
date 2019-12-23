package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"lokucrazy/Go-DnD-Info/models"
)

type proficienciesService struct {
	DB *sql.DB
}

func (p *proficienciesService) Get(name string) ([]byte, error) {
	if p.DB == nil {
		return nil, errors.New("db cannot be empty")
	}

	profs := make([]models.Proficiencies, 0)
	query := "SELECT * FROM proficiencies "
	if name != "" {
		query += "WHERE proficiencies.name = ?"
	}

	result, err := p.DB.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer closeRows(result)

	for result.Next() {
		prof := models.Proficiencies{}
		err = result.Scan(
			&prof.ID,
			&prof.Name,
			&prof.Class)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		profs = append(profs, prof)
	}

	if len(profs) == 1 {
		return json.Marshal(profs[0])
	}
	return json.Marshal(profs)
}

func CreateProficienciesService(db *sql.DB) *proficienciesService {
	return &proficienciesService{DB: db}
}
