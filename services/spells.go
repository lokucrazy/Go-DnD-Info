package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"lokucrazy/Go-DnD-Info/models"
)

type spellsService struct {
	DB *sql.DB
}

func (s *spellsService) Get(name string) ([]byte, error) {
	if s.DB == nil {
		return nil, errors.New("db cannot be nil")
	}

	var spells []models.Spells
	query := "SELECT * FROM spells "
	if name != "" {
		query += "WHERE retrievedSpells.name = ?"
	}

	result, err := s.DB.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer closeRows(result)

	for result.Next() {
		spell := models.Spells{}
		err = result.Scan(
			&spell.ID,
			&spell.Name,
			&spell.School,
			&spell.Level,
			&spell.Class,
			&spell.CastTime,
			&spell.Range,
			&spell.Components,
			&spell.Duration,
			&spell.Description)
		if err != nil {
			return nil, err
		}
		spells = append(spells, spell)
	}

	return json.Marshal(spells)
}

func CreateSpellsService(db *sql.DB) *spellsService {
	return &spellsService{DB: db}
}
