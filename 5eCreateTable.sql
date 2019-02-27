DROP TABLE IF EXISTS classes;
DROP TABLE IF EXISTS spells;
DROP TABLE IF EXISTS weapons;
DROP TABLE IF EXISTS armors;
DROP TABLE IF EXISTS races;
DROP TABLE IF EXISTS feats;
DROP TABLE IF EXISTS levels;
DROP TABLE IF EXISTS proficiencies;
CREATE TABLE classes (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    hitDie TEXT NOT NULL,
    startEquipment TEXT NOT NULL
);

CREATE TABLE spells (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    school TEXT NOT NULL,
    level INTEGER NOT NULL,
    class INTEGER NOT NULL,
    castTime INT NOT NULL,
    range TEXT NOT NULL,
    components TEXT NOT NULL,
    duration INT NOT NULL,
    description TEXT NOT NULL,
    FOREIGN KEY (class) REFERENCES classes(id)
);

CREATE TABLE weapons (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    damage TEXT NOT NULL,
    type TEXT NOT NULL,
    cost INTEGER NOT NULL,
    range TEXT NOT NULL,
    weight INTEGER NOT NULL,
    properties TEXT NOT NULL
);

CREATE TABLE armors (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    armorClass TEXT NOT NULL,
    type TEXT NOT NULL,
    cost INTEGER NOT NULL,
    strength TEXT NOT NULL,
    weight INTEGER NOT NULL,
    stealth TEXT NOT NULL
);

CREATE TABLE races (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE feats (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    preRequisite TEXT NOT NULL,
    description TEXT NOT NULL,
    class INTEGER NOT NULL,
    level INTEGER NOT NULL,
    FOREIGN KEY(class) REFERENCES classes(id)
);

CREATE TABLE levels (
    id INTEGER PRIMARY KEY,
    proficiencyBonus INTEGER NOT NULL,
    class INTEGER NOT NULL,
    FOREIGN KEY(class) REFERENCES classes(id)
);

CREATE TABLE proficiencies (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    class INTEGER NOT NULL,
    FOREIGN KEY(class) REFERENCES classes(id)
);

INSERT INTO classes VALUES (1,'Bard','1d8','(a) a rapier, (b) a longsword, or (c) any simple weapon
(a) a diplomat’s pack or (b) an entertainer’s pack
(a) a lute or (b) any other musical instrument
Leather armor and a dagger');
INSERT INTO spells VALUES (1,'Blade Ward','Abjuration',0,1,1,'self','V,S',1,'Tracing an arcane sigil in the air, you grant yourself resistance to bludgeoning, piercing, and slashing damage dealt by weapons until the end of your next turn.');
INSERT INTO races VALUES(1,'Dragonborn');
INSERT INTO races VALUES(2,'Human');
INSERT INTO races VALUES(3,'Elf');