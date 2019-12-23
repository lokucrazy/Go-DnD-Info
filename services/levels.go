package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"lokucrazy/Go-DnD-Info/models"
)

type levelsService struct {
	DB *sql.DB
}

func (l *levelsService) Get(name string) ([]byte, error) {
	if l.DB == nil {
		return nil, errors.New("db cannot be nil")
	}

	levels := make([]models.Levels, 0)
	query := "SELECT * FROM levels "
	if name != "" {
		query += "WHERE levels.name = ?"
	}

	result, err := l.DB.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer closeRows(result)

	for result.Next() {
		level := models.Levels{}
		err = result.Scan(
			&level.ID,
			&level.ProficiencyBonus,
			&level.Class)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		levels = append(levels, level)
	}

	if len(levels) == 1 {
		return json.Marshal(levels[0])
	}
	return json.Marshal(levels)
}

func CreateLevelsService(db *sql.DB) *levelsService {
	return &levelsService{DB: db}
}
