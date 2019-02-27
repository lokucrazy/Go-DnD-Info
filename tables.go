package main

//Classes Struct
type Classes struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	HitDie         string `json:"hitDie,omitempty"`
	StartEquipment string `json:"startEquipment,omitempty"`
}

//Spells struct
type Spells struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	School      string `json:"school,omitempty"`
	Level       int    `json:"level,omitempty"`
	Class       int    `json:"class,omitempty"`
	CastTime    int    `json:"castTime,omitempty"`
	Range       string `jsong:"range,omitempty"`
	Components  string `json:"components,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Description string `json:"description,omitempty"`
}

//Weapons struct
type Weapons struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Damage     string `json:"damage,omitempty"`
	Type       string `json:"type,omitempty"`
	Cost       int    `json:"cost,omitempty"`
	Range      string `json:"range,omitempty"`
	Weight     int    `json:"weight,omitempty"`
	Properties string `json:"properties,omitempty"`
}

//Armor struct
type Armors struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	ArmorClass string `json:"armorClass,omitempty"`
	Type       string `json:"type,omitempty"`
	Cost       int    `json:"cost,omitempty"`
	Strength   string `json:"strength,omitempty"`
	Weight     int    `json:"weight,omitempty"`
	Stealth    string `json:"stealth,omitempty"`
}

//Feats struct
type Feats struct {
	ID           int    `json:"id,omitempty"`
	PreRequisite string `json:"preRequisite,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Class        int    `json:"class,omitempty"`
	Level        int    `json:"level,omitempty"`
}

//Levels struct
type Levels struct {
	ID               int `json:"id,omitempty"`
	ProficiencyBonus int `json:"proficiencyBonus,omitempty"`
	Class            int `json:"class,omitempty"`
}

//Proficiencies struct
type Proficiencies struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Class int    `json:"class,omitempty"`
}

//Races struct
type Races struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
