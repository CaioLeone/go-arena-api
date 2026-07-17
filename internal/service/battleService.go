package service

import (
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/battle"
	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/game"
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
	characterService   CharacterService
	leaderboardService *ranking.LeaderboardService
}

func NewBattleService(battleRepo repository.BattleRepository, characterRepo repository.CharacterRepository, characterServ CharacterService, leaderboardService *ranking.LeaderboardService) BattleService {
	return &battleService{
		battleRepo:         battleRepo,
		characterRepo:      characterRepo,
		characterService:   characterServ,
		leaderboardService: leaderboardService,
	}
}

func (s *battleService) loadCharacters(userID string, req *dto.BattleCreateRequest) (*model.CharacterModel, *model.CharacterModel, error) {
	//Busca atacante
	attacker, err := s.characterRepo.GetByID(req.AttackerCharacterID, userID)
	if err != nil {
		return nil, nil, fmt.Errorf("Personagem atacante não encontrado")
	}

	//Busca Defensor(qualquer usuario)
	defender, err := s.characterRepo.GetByIDNoUserFilter(req.DefenderCharacterID)
	if err != nil {
		return nil, nil, fmt.Errorf("Personagem defensor não encontrado")
	}

	return attacker, defender, nil
}

func (s *battleService) buildBattleModel(attacker *model.CharacterModel, defender *model.CharacterModel, battleResult *battle.BattleResult) (*model.BattleModel, error) {
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
	}

	attackerDamage := attacker.HP - battleResult.AttackerHPFinal
	if attackerDamage < 0 {
		attackerDamage = 0
	}

	defenderDamage := defender.HP - battleResult.DefenderHPFinal
	if defenderDamage < 0 {
		defenderDamage = 0
	}

	damageDealt := attackerDamage + defenderDamage
	if damageDealt < 0 {
		damageDealt = 0
	}

	//Salvar Batalha no Banco
	return &model.BattleModel{
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
	}, nil
}

func (s *battleService) updateRanking(result *battle.BattleResult) {
	if result.IsDraw || result.Winner == nil {
		return
	}

	winnerRanking, _ := s.battleRepo.GetCharacterRanking(result.Winner.ID.String())
	loserRanking, _ := s.battleRepo.GetCharacterRanking(result.Loser.ID.String())
	rankingDiff := winnerRanking - loserRanking

	pointsGained := battle.UpdateRanking(result.Winner, result.Loser, rankingDiff)

	newWinnerScore := winnerRanking + pointsGained
	newLoserScore := loserRanking - 5

	if newLoserScore < 0 {
		newLoserScore = 0
	}

	// Atualizar no banco
	s.battleRepo.UpdateCharacterRanking(result.Winner.ID.String(), pointsGained)
	s.battleRepo.UpdateCharacterRanking(result.Loser.ID.String(), newLoserScore-loserRanking)

	s.leaderboardService.UpdatePlayerScore(
		result.Winner.ID.String(),
		result.Winner.Name,
		float64(newWinnerScore),
	)
	s.leaderboardService.UpdatePlayerScore(
		result.Loser.ID.String(),
		result.Loser.Name,
		float64(newLoserScore),
	)
}

func (s *battleService) buildBattleResponse(battleModel *model.BattleModel) (*dto.BattleResponse, error) {
	rounds, err := battle.FromRoundsJSON(battleModel.RoundsData)
	if err != nil {
		return nil, fmt.Errorf("Erro ao desserializar rounds: %w", err)
	}

	return &dto.BattleResponse{
		ID:              battleModel.ID,
		AttackerID:      battleModel.AttackerID,
		AttackerName:    battleModel.AttackerName,
		DefenderID:      battleModel.DefenderID,
		DefenderName:    battleModel.DefenderName,
		WinnerID:        battleModel.WinnerID,
		WinnerName:      battleModel.WinnerName,
		DamageDealt:     battleModel.DamageDealt,
		AttackerHPFinal: battleModel.AttackerHPFinal,
		DefenderHPFinal: battleModel.DefenderHPFinal,
		RoundsCount:     battleModel.RoundsCount,
		Rounds:          rounds,
		CreatedAt:       battleModel.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *battleService) StartBattle(userID string, req *dto.BattleCreateRequest) (*dto.BattleResponse, error) {

	//Carrega Jogadores
	attacker, defender, err := s.loadCharacters(userID, req)
	if err != nil {
		return nil, err
	}

	//Simular Batalha
	battleResult, err := battle.DetermineBattle(attacker, defender)
	if err != nil {
		return nil, fmt.Errorf("Erro ao Simular Batalha: %w", err)
	}

	battleModel, err := s.buildBattleModel(attacker, defender, battleResult)
	if err != nil {
		return nil, fmt.Errorf("Erro ao Construir Modelo de Batalha: %w", err)
	}

	savedBattle, err := s.battleRepo.Create(battleModel)
	if err != nil {
		return nil, fmt.Errorf("Erro ao Salvar Batalha: %w", err)
	}

	s.updateRanking(battleResult)

	if !battleResult.IsDraw {
		s.characterService.AddExperience(battleResult.Winner.ID.String(), game.WinnerExperience)
		s.characterService.AddExperience(battleResult.Loser.ID.String(), game.LoserExperience)
	}

	return s.buildBattleResponse(savedBattle)
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
