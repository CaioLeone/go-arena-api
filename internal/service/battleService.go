package service

import (
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/battle"
	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/model"
	"github.com/caioLeone/go-arena-api/internal/repository"
	"github.com/google/uuid"
)

type BattleService interface {
	StartBattle(userID string, req *dto.BattleCreateRequest) (*dto.BattleResponse, error)
	GetBattleHistory(userID string, limit int, offset int) ([]*dto.BattleHistoryResponse, error)
	GetBattleByID(battleID string) (*dto.BattleResponse, error)
}

type battleService struct {
	battleRepo    repository.BattleRepository
	characterRepo repository.CharacterRepository
}

func NewBattleService(battleRepo repository.BattleRepository, characterRepo repository.CharacterRepository) BattleService {
	return &battleService{
		battleRepo:    battleRepo,
		characterRepo: characterRepo,
	}
}

func (s *battleService) StartBattle(userID string, req *dto.BattleCreateRequest) (*dto.BattleResponse, error) {
	//Busca atacante
	attacker, err := s.characterRepo.GetByID(req.AttackerCharacterID, userID)
	if err != nil {
		return nil, fmt.Errorf("Personagem Atacante Nao Encontrado")
	}

	//Busca Defensor(qualquer usuario)
	defender, err := s.characterRepo.GetByID(req.DefenderCharacterID, "")
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

	winnerID := uuid.Nil
	winnerName := "Empate"
	if battleResult.Winner != nil {
		winnerID = battleResult.Winner.ID
		winnerName = battleResult.Winner.Name
	}

	//Salvar Batalha no Banco
	battleModel := &model.BattleModel{
		AttackerID:      attacker.ID,
		AttackerName:    attacker.Name,
		DefenderID:      defender.ID,
		DefenderName:    defender.Name,
		WinnerID:        winnerID,
		WinnerName:      winnerName,
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
		winnerRanking, _ := s.battleRepo.GetCharacterRanking(string(battleResult.Winner.ID.String()))
		loserRanking, _ := s.battleRepo.GetCharacterRanking(string(battleResult.Loser.ID.String()))

		RankingDiff := winnerRanking - loserRanking
		pointsGained := battle.UpdateRanking(battleResult.Winner, battleResult.Loser, RankingDiff)

		s.battleRepo.UpdateCharacterRanking(battleResult.Winner.ID.String(), pointsGained)
		s.battleRepo.UpdateCharacterRanking(battleResult.Loser.ID.String(), -5) //Perde 5 Pontos
	}

	rounds, _ := battle.FromRoundsJSON(roundsJSON)
	return &dto.BattleResponse{
		ID:              savedBattle.ID,
		AttackerID:      savedBattle.AttackerID,
		AttackerName:    savedBattle.AttackerName,
		DefenderID:      savedBattle.DefenderID,
		DefenderName:    savedBattle.DefenderName,
		WinnerID:        savedBattle.WinnerID,
		WinnerName:      savedBattle.WinnerName,
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
		AttackerHPFinal: b.AttackerHPFinal,
		DefenderHPFinal: b.DefenderHPFinal,
		RoundsCount:     b.RoundsCount,
		Rounds:          rounds,
		CreatedAt:       b.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
