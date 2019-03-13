#Go-DnD

DnD 5e API in Go

example query:

    http://url:8080/get/weapons
    
    http://url:8080/get/races/human
    
example response:

    [{
        "id": 1,
        "name": "club",
        "damage": "1d4 bludgeoning",
        "type": "simple weapon",
        "cost": "1 sp",
        "range": null,
        "weight": "2 lb",
        "properties": "light"
    }]

you can either not specify a item, and that will send a list of all the items.
Or you can specify what item you want, and it will send a list of size 1 with that item.

Spells:

    {
        id
        name
        school
        level
        class
        castTime
        range
        components
        duration
        description
    }
    
Weapons:

    {
        id
        name
        damage
        type
        cost
        range
        weight
        properties
    }
    
Armor:
    
    {
    	id
    	name
    	type
    	cost
    	armorClass
    	strength
    	stealth
    	weight
    }
    
Races:
    
    {
    	id
    	name
    }
