DROP TABLE IF EXISTS spells;
DROP TABLE IF EXISTS weapons;
DROP TABLE IF EXISTS armors;
DROP TABLE IF EXISTS races;
DROP TABLE IF EXISTS feats;
DROP TABLE IF EXISTS levels;
DROP TABLE IF EXISTS proficiencies;

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
    description TEXT NOT NULL
);

CREATE TABLE weapons (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    damage TEXT,
    type TEXT NOT NULL,
    cost TEXT NOT NULL,
    range TEXT,
    weight TEXT,
    properties TEXT
);

CREATE TABLE armors (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    cost TEXT NOT NULL,
    armorClass TEXT NOT NULL,
    strength TEXT,
    stealth TEXT,
    weight TEXT NOT NULL
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
    level INTEGER NOT NULL
);

CREATE TABLE levels (
    id INTEGER PRIMARY KEY,
    proficiencyBonus INTEGER NOT NULL,
    class INTEGER NOT NULL
);

CREATE TABLE proficiencies (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    class INTEGER NOT NULL
);

INSERT INTO spells VALUES (1,'blade ward','Abjuration',0,1,1,'self','V,S',1,'Tracing an arcane sigil in the air, you grant yourself resistance to bludgeoning, piercing, and slashing damage dealt by weapons until the end of your next turn.');
INSERT INTO races VALUES(1,'dragonborn');
INSERT INTO races VALUES(2,'human');
INSERT INTO races VALUES(3,'elf');

INSERT INTO armors VALUES(1,'padded','light armor','5 gp','11+dex modifier',null,'disadvantage','8 lb'),
                         (2,'leather','light armor','10 gp','11+dex modifier',null,null,'10 lb'),
                         (3,'studded leather','light armor','45 gp','12+dex modifier',null,null,'13 lb'),
                         (4,'hide','medium armor','10 gp','12+dex modifier(max 2)',null,null,'12 lb'),
                         (5,'chain shirt','medium armor', '50 gp','13+dex modifier(max 2)',null,null,'20 lb'),
                         (6,'scale mail','medium armor','50 gp','14+dex modifier(max 2)',null,'disadvantage','45 lb'),
                         (7,'breastplate','medium armor','400 gp','14+dex modifier(max 2)',null,null,'20 lb'),
                         (9,'half plate','medium armor','750 gp','15+dex modifier(max 2)',null,'disadvantage','40 lb'),
                         (10,'ring mail','heavy armor','30 gp','14',null,'disadvantage','40 lb'),
                         (11,'chain mail','heavy armor','75 gp','16','str 13','disadvantage','55 lb'),
                         (12,'splint','heavy armor','200 gp','17','str 15','disadvantage','60 lb'),
                         (13,'plate','heavy armor','1,500 gp','18','str 15','disadvantage','65 lb'),
                         (14,'shield','shield','10 gp','+2',null,null,'6 lb');

INSERT INTO weapons VALUES(1,'club','1d4 bludgeoning','simple weapon','1 sp',null,'2 lb','light'),
                          (2,'dagger','1d4 piercing','simple weapon','1 gp','20/60','1 lb','finesse,light,thrown'),
                          (3,'greatclub','1d8 bludgeoning','simple weapon','2 sp',null,'10 lb','two-handed'),
                          (4,'handaxe','1d6 slashing','simple weapon','5 gp','20/60','2 lb','light,thrown'),
                          (5,'javelin','1d6 piercing','simple weapon','5 sp','31/120','2 lb','thrown'),
                          (6,'light hammer','1d4 bludgeoning','simple weapon','2 gp','20/60','2 lb','light,thrown'),
                          (7,'mace','1d6 bludgeoning','simple weapon','5 gp',null,'4 lb',null),
                          (8,'quarterstaff','1d6 bludgeoning','simple weapon','2 sp',null,'4 lb','Versatile(1d8)'),
                          (9,'sickle','1d5 slashing','simple weapon','1 gp',null,'2 lb','light'),
                          (10,'spear','1d6 piercing','simple weapon','1 gp','20/60','3 lb','thrown,versatile(1d8)');