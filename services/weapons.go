package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"lokucrazy/Go-DnD-Info/models"
)

type weaponsService struct {
	DB *sql.DB
}

func (w *weaponsService) Get(name string) ([]byte, error) {
	if w.DB == nil {
		return nil, errors.New("db cannot be nil")
	}

	weapons := make([]models.Weapons, 0)
	query := "SELECT * FROM weapons "
	if name != "" {
		query += "WHERE weapons.name = ?"
	}

	result, err := w.DB.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer closeRows(result)

	for result.Next() {
		weapon := models.Weapons{}
		err = result.Scan(
			&weapon.ID,
			&weapon.Name,
			&weapon.Damage,
			&weapon.Type,
			&weapon.Cost,
			&weapon.Range,
			&weapon.Weight,
			&weapon.Properties)
		if err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}

	if len(weapons) == 1 {
		return json.Marshal(weapons[0])
	}
	return json.Marshal(weapons)
}

func CreateWeaponsService(db *sql.DB) *weaponsService {
	return &weaponsService{DB: db}
}
