package models

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

//Spells struct
type Spells struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	School      string `json:"school"`
	Level       int    `json:"level"`
	Class       int    `json:"class"`
	CastTime    int    `json:"castTime"`
	Range       string `json:"range"`
	Components  string `json:"components"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
}

//Weapons struct
type Weapons struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Damage     *NullString `json:"damage"`
	Type       string      `json:"type"`
	Cost       string      `json:"cost"`
	Range      *NullString `json:"range"`
	Weight     *NullString `json:"weight"`
	Properties *NullString `json:"properties"`
}

//Armor struct
type Armors struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Cost       string      `json:"cost"`
	ArmorClass string      `json:"armorClass"`
	Strength   *NullString `json:"strength"`
	Stealth    *NullString `json:"stealth"`
	Weight     string      `json:"weight"`
}

//Races struct
type Races struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Feats struct
type Feats struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PreRequisite string `json:"preRequisite"`
	Description  string `json:"description"`
	Class        int    `json:"class"`
	Level        int    `json:"level"`
}

//Levels struct
type Levels struct {
	ID               int `json:"id"`
	ProficiencyBonus int `json:"proficiencyBonus"`
	Class            int `json:"class"`
}

//Proficiencies struct
type Proficiencies struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class int    `json:"class"`
}

type NullString sql.NullString

func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ns = NullString{String: s.String, Valid: false}
	} else {
		*ns = NullString{String: s.String, Valid: true}
	}
	return nil
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return nil, nil
	}
	return json.Marshal(ns.String)
}
