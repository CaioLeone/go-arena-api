package battle

import (
	"math"
	"math/rand"
	"time"

	"github.com/caioLeone/go-arena-api/internal/model"
)

const (
	MaxRounds = 50
	MinDamage = 1
)

type BattleRoundData struct {
	AttackerDamage int    `json:"attacker_damage"`
	DefenderHP     int    `json:"defender_hp"`
	IsCritical     bool   `json:"is_critical"`
	Message        string `json:"message"`
}

//CalcDamage Calcula o dano usando a formula especificada
// Dano = (Ataque × (100 / (100 + Defesa))) × Random(0.9, 1.1) × Crítico

func CalcDamage(attacker *model.CharacterModel, defender *model.CharacterModel) (int, bool) {
	rand.Seed(time.Now().UnixNano())

	baseDamage := float64(attacker.Attack) * (100.0 / (100.0 + float64(defender.Defense)))

	randomMultiplier := 0.9 + rand.Float64()*(1.1-0.9)
	damage := baseDamage * randomMultiplier

	isCritical := false
	if rand.Intn(100) < attacker.CriticalChance {
		damage *= 1.5
		isCritical = true
	}

	finalDamage := int(math.Max(float64(MinDamage), damage))

	return finalDamage, isCritical
}
