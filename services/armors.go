package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"lokucrazy/Go-DnD-Info/models"
)

type armorsService struct {
	DB *sql.DB
}

func (a *armorsService) Get(name string) ([]byte, error) {
	if a.DB == nil {
		return nil, errors.New("db cannot be nil")
	}

	armors := make([]models.Armors, 0)
	query := "SELECT * FROM armors "
	if name != "" {
		query += "WHERE armors.name = ?"
	}

	result, err := a.DB.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer closeRows(result)

	for result.Next() {
		armor := models.Armors{}
		err = result.Scan(
			&armor.ID,
			&armor.Name,
			&armor.Type,
			&armor.Cost,
			&armor.ArmorClass,
			&armor.Strength,
			&armor.Stealth,
			&armor.Weight)
		if err != nil {
			return nil, err
		}
		armors = append(armors, armor)
	}

	if len(armors) == 1 {
		return json.Marshal(armors[0])
	}
	return json.Marshal(armors)
}

func CreateArmorsService(db *sql.DB) *armorsService {
	return &armorsService{DB: db}
}
