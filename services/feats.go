package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"lokucrazy/Go-DnD-Info/models"
)

type featsService struct {
	DB *sql.DB
}

func (f *featsService) Get(name string) ([]byte, error) {
	if f.DB == nil {
		return nil, errors.New("db cannot be nil")
	}

	feats := make([]models.Feats, 0)
	query := "SELECT * FROM feats "
	if name != "" {
		query += "WHERE feats.name = ?"
	}

	result, err := f.DB.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer closeRows(result)

	for result.Next() {
		feat := models.Feats{}
		err = result.Scan(
			&feat.ID,
			&feat.Name,
			&feat.PreRequisite,
			&feat.Description,
			&feat.Class,
			&feat.Level)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		feats = append(feats, feat)
	}

	if len(feats) == 1 {
		return json.Marshal(feats[0])
	}
	return json.Marshal(feats)
}

func CreateFeatsService(db *sql.DB) *featsService {
	return &featsService{DB: db}
}
