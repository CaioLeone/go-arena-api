package service

import (
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/battle"
	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/model"
	"github.com/caioLeone/go-arena-api/internal/ranking"
	"github.com/caioLeone/go-arena-api/internal/repository"
	"github.com/google/uuid"
)

type BattleService interface {
	StartBattle(userID string, req *dto.BattleCreateRequest) (*dto.BattleResponse, error)
	GetBattleHistory(userID string, limit int, offset int) ([]*dto.BattleHistoryResponse, error)
	GetBattleByID(battleID string) (*dto.BattleResponse, error)
}

type battleService struct {
	battleRepo         repository.BattleRepository
	characterRepo      repository.CharacterRepository
	leaderboardService *ranking.LeaderboardService
}

func NewBattleService(battleRepo repository.BattleRepository, characterRepo repository.CharacterRepository, leaderboardService *ranking.LeaderboardService) BattleService {
	return &battleService{
		battleRepo:         battleRepo,
		characterRepo:      characterRepo,
		leaderboardService: leaderboardService,
	}
}

func (s *battleService) StartBattle(userID string, req *dto.BattleCreateRequest) (*dto.BattleResponse, error) {
	//Busca atacante
	attacker, err := s.characterRepo.GetByID(req.AttackerCharacterID, userID)
	if err != nil {
		return nil, fmt.Errorf("Personagem atacante não encontrado")
	}

	//Busca Defensor(qualquer usuario)
	defender, err := s.characterRepo.GetByIDNoUserFilter(req.DefenderCharacterID)
	if err != nil {
		return nil, fmt.Errorf("Personagem defensor não encontrado")
	}

	//Simular Batalha
	battleResult, err := battle.DetermineBattle(attacker, defender)
	if err != nil {
		return nil, fmt.Errorf("Erro ao Simular Batalha: %w", err)
	}

	//Serializar Rounds
	roundsJSON, err := battleResult.ToRoundsJSON()
	if err != nil {
		return nil, fmt.Errorf("Erro ao Serializar Rounds: %w", err)
	}

	var winnerID *uuid.UUID
	winnerName := "Empate"
	if battleResult.Winner != nil {
		winnerID = &battleResult.Winner.ID
		winnerName = battleResult.Winner.Name
	} else {
		winnerID = nil
		winnerName = "Empate"
	}

	damageDealt := defender.HP - battleResult.DefenderHPFinal
	if damageDealt < 0 {
		damageDealt = 0
	}

	//Salvar Batalha no Banco
	battleModel := &model.BattleModel{
		AttackerID:      attacker.ID,
		AttackerName:    attacker.Name,
		DefenderID:      defender.ID,
		DefenderName:    defender.Name,
		WinnerID:        winnerID,
		WinnerName:      winnerName,
		DamageDealt:     damageDealt,
		AttackerHPFinal: battleResult.AttackerHPFinal,
		DefenderHPFinal: battleResult.DefenderHPFinal,
		RoundsCount:     battleResult.RoundsCount,
		RoundsData:      roundsJSON,
	}

	savedBattle, err := s.battleRepo.Create(battleModel)
	if err != nil {
		return nil, fmt.Errorf("Erro ao Salvar Batalha: %w", err)
	}

	//Atualizar Ranking se nao foi empate
	if !battleResult.IsDraw && battleResult.Winner != nil {
		winnerRanking, _ := s.battleRepo.GetCharacterRanking(battleResult.Winner.ID.String())
		loserRanking, _ := s.battleRepo.GetCharacterRanking(battleResult.Loser.ID.String())

		rankingDiff := winnerRanking - loserRanking
		pointsGained := battle.UpdateRanking(battleResult.Winner, battleResult.Loser, rankingDiff)

		// Atualizar no banco
		s.battleRepo.UpdateCharacterRanking(battleResult.Winner.ID.String(), pointsGained)
		s.battleRepo.UpdateCharacterRanking(battleResult.Loser.ID.String(), -5)

		//ADICIONADO: Atualizar no Redis Leaderboard
		newWinnerScore := float64(winnerRanking + pointsGained)
		newLoserScore := float64(loserRanking - 5)

		s.leaderboardService.UpdatePlayerScore(
			battleResult.Winner.ID.String(),
			battleResult.Winner.Name,
			newWinnerScore,
		)
		s.leaderboardService.UpdatePlayerScore(
			battleResult.Loser.ID.String(),
			battleResult.Loser.Name,
			newLoserScore,
		)
	}

	//rounds, _ := battle.FromRoundsJSON(roundsJSON)
	rounds, err := battle.FromRoundsJSON(savedBattle.RoundsData)
	if err != nil {
		return nil, fmt.Errorf("Erro ao desserializar rounds: %w", err)
	}
	return &dto.BattleResponse{
		ID:              savedBattle.ID,
		AttackerID:      savedBattle.AttackerID,
		AttackerName:    savedBattle.AttackerName,
		DefenderID:      savedBattle.DefenderID,
		DefenderName:    savedBattle.DefenderName,
		WinnerID:        savedBattle.WinnerID,
		WinnerName:      savedBattle.WinnerName,
		DamageDealt:     savedBattle.DamageDealt,
		AttackerHPFinal: savedBattle.AttackerHPFinal,
		DefenderHPFinal: savedBattle.DefenderHPFinal,
		RoundsCount:     savedBattle.RoundsCount,
		Rounds:          rounds,
		CreatedAt:       savedBattle.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *battleService) GetBattleHistory(userID string, limit int, offset int) ([]*dto.BattleHistoryResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	battles, err := s.battleRepo.GetHistoryByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []*dto.BattleHistoryResponse
	for _, b := range battles {
		responses = append(responses, &dto.BattleHistoryResponse{
			ID:              b.ID,
			AttackerID:      b.AttackerID,
			AttackerName:    b.AttackerName,
			DefenderID:      b.DefenderID,
			DefenderName:    b.DefenderName,
			WinnerID:        b.WinnerID,
			WinnerName:      b.WinnerName,
			DamageDealt:     b.DamageDealt,
			AttackerHPFinal: b.AttackerHPFinal,
			DefenderHPFinal: b.DefenderHPFinal,
			RoundsCount:     b.RoundsCount,
			CreatedAt:       b.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return responses, nil
}

func (s *battleService) GetBattleByID(battleID string) (*dto.BattleResponse, error) {
	b, err := s.battleRepo.GetByID(battleID)
	if err != nil {
		return nil, err
	}

	rounds, _ := battle.FromRoundsJSON(b.RoundsData)

	return &dto.BattleResponse{
		ID:              b.ID,
		AttackerID:      b.AttackerID,
		AttackerName:    b.AttackerName,
		DefenderID:      b.DefenderID,
		DefenderName:    b.DefenderName,
		WinnerID:        b.WinnerID,
		WinnerName:      b.WinnerName,
		DamageDealt:     b.DamageDealt,
		AttackerHPFinal: b.AttackerHPFinal,
		DefenderHPFinal: b.DefenderHPFinal,
		RoundsCount:     b.RoundsCount,
		Rounds:          rounds,
		CreatedAt:       b.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
