package game

type CharacterStats struct {
	HP             int
	Attack         int
	Defense        int
	CriticalChance int
}

var InitialStats = map[string]CharacterStats{
	"Barbaro": {
		HP:             120,
		Attack:         20,
		Defense:        8,
		CriticalChance: 10,
	},
	"Mago": {
		HP:             80,
		Attack:         12,
		Defense:        5,
		CriticalChance: 15,
	},
	"Arqueiro": {
		HP:             90,
		Attack:         14,
		Defense:        6,
		CriticalChance: 20,
	},
	"Assassino": {
		HP:             85,
		Attack:         18,
		Defense:        4,
		CriticalChance: 25,
	},
}
