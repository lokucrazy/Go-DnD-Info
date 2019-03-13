package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func get(table, name string) ([]byte, error) {
	db, err := sql.Open("sqlite3", "5eInfo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var list interface{}
	switch table {
	case "spells":
		list, err = getSpells(db, name)
	case "weapons":
		list, err = getWeapons(db, name)
	case "armors":
		list, err = getArmors(db, name)
	case "races":
		list, err = getRaces(db, name)
	case "feats":
		list, err = getFeats(db, name)
	case "levels":
		list, err = getLevels(db, name)
	case "proficiencies":
		list, err = getProficiencies(db, name)
	}
	if err != nil {
		return nil, err
	}

	return json.Marshal(list)
}

func getSpells(db *sql.DB, name string) ([]Spells, error) {
	var spells []Spells
	query := "SELECT * FROM spells "
	if name != "" {
		query += "WHERE spells.name = ?"
	}
	result, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		spell := Spells{}
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
	return spells, nil
}

func getWeapons(db *sql.DB, name string) ([]Weapons, error) {
	var weapons []Weapons
	query := "SELECT * FROM weapons "
	if name != "" {
		query += "WHERE weapons.name = ?"
	}
	result, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		weapon := Weapons{}
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
	return weapons, nil
}

func getArmors(db *sql.DB, name string) ([]Armors, error) {
	var armors []Armors
	query := "SELECT * FROM armors "
	if name != "" {
		query += "WHERE armors.name = ?"
	}
	result, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		armor := Armors{}
		err = result.Scan(
			&armor.ID,
			&armor.Name,
			&armor.ArmorClass,
			&armor.Type,
			&armor.Cost,
			&armor.Strength,
			&armor.Weight,
			&armor.Stealth)
		if err != nil {
			return nil, err
		}
		armors = append(armors, armor)
	}
	return armors, nil
}

func getRaces(db *sql.DB, name string) ([]Races, error) {
	var races []Races
	query := "SELECT * FROM races "
	if name != "" {
		query += "WHERE races.name = ?"
	}
	result, err := db.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for result.Next() {
		race := Races{}
		err = result.Scan(
			&race.ID,
			&race.Name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		races = append(races, race)
	}
	return races, nil
}

func getFeats(db *sql.DB, name string) ([]Feats, error) {
	var feats []Feats
	query := "SELECT * FROM feats "
	if name != "" {
		query += "WHERE feats.name = ?"
	}
	result, err := db.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for result.Next() {
		feat := Feats{}
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
	return feats, nil
}

func getLevels(db *sql.DB, name string) ([]Levels, error) {
	var levels []Levels
	query := "SELECT * FROM levels "
	if name != "" {
		query += "WHERE levels.name = ?"
	}
	result, err := db.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for result.Next() {
		level := Levels{}
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
	return levels, nil
}

func getProficiencies(db *sql.DB, name string) ([]Proficiencies, error) {
	var profs []Proficiencies
	query := "SELECT * FROM proficiencies "
	if name != "" {
		query += "WHERE proficiencies.name = ?"
	}
	result, err := db.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for result.Next() {
		prof := Proficiencies{}
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
	return profs, nil
}
