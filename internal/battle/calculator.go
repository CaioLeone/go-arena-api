package battle

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/model"
)

const (
	MaxRounds = 50
	MinDamage = 1
)

type BattleRoundData struct {
	Round        int    `json:"round"`
	AttackerName string `json:"attacker_name"`
	DefenderName string `json:"defender_name"`
	Damage       int    `json:"damage"`
	RemainingHP  int    `json:"remaining_hp"`
	IsCritical   bool   `json:"is_critical"`
	Message      string `json:"message"`
}

type BattleResult struct {
	Winner          *model.CharacterModel
	Loser           *model.CharacterModel
	AttackerHPFinal int
	DefenderHPFinal int
	RoundsCount     int
	Rounds          []BattleRoundData
	IsDraw          bool
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

// DetermineBattle simula a batalha entre dois personagens e retorna o resultado
func DetermineBattle(attacker, defender *model.CharacterModel) (*BattleResult, error) {
	if attacker.ID == defender.ID {
		return nil, fmt.Errorf("Um Personagem Nao Pode Lutar Contra Si Mesmo")
	}

	rounds := []BattleRoundData{}
	attackerHP := attacker.HP
	defenderHP := defender.HP

	//Simular Rounds de Batalha
	for round := 0; round < MaxRounds; round++ {
		//-------------------------
		//Atacante Ataca
		//-------------------------
		damage, isCritical := CalcDamage(attacker, defender)
		defenderHP -= damage

		if defenderHP < 0 {
			defenderHP = 0
		}

		rounds = append(rounds, BattleRoundData{
			Round:        round + 1,
			AttackerName: attacker.Name,
			DefenderName: defender.Name,
			Damage:       damage,
			RemainingHP:  defenderHP,
			IsCritical:   isCritical,
			Message:      fmt.Sprintf("%s causou %d de dano em %s", attacker.Name, damage, defender.Name),
		})

		if defenderHP == 0 {
			return &BattleResult{
				Winner:          attacker,
				Loser:           defender,
				AttackerHPFinal: attackerHP,
				DefenderHPFinal: 0,
				RoundsCount:     round + 1,
				Rounds:          rounds,
				IsDraw:          false,
			}, nil
		}

		//-------------------------
		//Defender Contra-Ataca
		//-------------------------

		damage, isCritical = CalcDamage(defender, attacker)
		attackerHP -= damage
		if attackerHP < 0 {
			attackerHP = 0
		}
		rounds = append(rounds, BattleRoundData{
			Round:        round + 1,
			AttackerName: defender.Name,
			DefenderName: attacker.Name,
			Damage:       damage,
			RemainingHP:  attackerHP,
			IsCritical:   isCritical,
			Message:      fmt.Sprintf("%s causou %d de dano em %s", defender.Name, damage, attacker.Name),
		})

		if attackerHP == 0 {
			return &BattleResult{
				Winner:          defender,
				Loser:           attacker,
				AttackerHPFinal: 0,
				DefenderHPFinal: defenderHP,
				RoundsCount:     round + 1,
				Rounds:          rounds,
				IsDraw:          false,
			}, nil
		}
	}

	//Se Chegou aqui, Empatou
	return &BattleResult{
		Winner:          nil,
		Loser:           nil,
		AttackerHPFinal: attackerHP,
		DefenderHPFinal: defenderHP,
		RoundsCount:     MaxRounds,
		Rounds:          rounds,
		IsDraw:          true,
	}, nil
}

func UpdateRanking(winner, loser *model.CharacterModel, diff int) int {
	pointsGained := 10 + (diff / 5)

	if pointsGained < 1 {
		pointsGained = 1
	}
	return pointsGained
}

// ToRoundsJSON Converte os dados dos rounds para JSON, facilitando a visualização e análise dos resultados da batalha.
func (br *BattleResult) ToRoundsJSON() (string, error) {
	data, err := json.Marshal(br.Rounds)
	if err != nil {
		return "", fmt.Errorf("Erro ao Serializar Rounds: %w", err)
	}
	return string(data), nil
}

// FromRoundsJSON desserialzia rounds de JSON
func FromRoundsJSON(data string) ([]dto.BattleRound, error) {
	var rounds []dto.BattleRound
	err := json.Unmarshal([]byte(data), &rounds)
	if err != nil {
		return nil, fmt.Errorf("Erro ao Desserializar Rounds: %w", err)
	}
	return rounds, nil
}
