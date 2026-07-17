package service

import (
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/game"
	"github.com/caioLeone/go-arena-api/internal/model"
	"github.com/caioLeone/go-arena-api/internal/repository"
	"github.com/google/uuid"
)

type CharacterService interface {
	Create(userID string, req *dto.CharacterCreateRequest) (*dto.CharacterResponse, error)
	GetByID(id string, userID string) (*dto.CharacterResponse, error)
	GetAll(userID string) ([]*dto.CharacterResponse, error)
	Update(id string, userID string, req *dto.CharacterUpdateRequest) (*dto.CharacterResponse, error)
	Delete(id string, userID string) error
	AddExperience(characterID string, experience int) error
	SpendAttributePoints(characterID string, req *dto.SpendAttributePointsRequest) error
}

type characterService struct {
	characterRepo repository.CharacterRepository
}

func NewCharacterService(characterRepo repository.CharacterRepository) CharacterService {
	return &characterService{characterRepo: characterRepo}
}

func modelToDTO(char *model.CharacterModel) *dto.CharacterResponse {
	return &dto.CharacterResponse{
		ID:              char.ID,
		UserID:          char.UserID,
		Name:            char.Name,
		Class:           char.Class,
		Level:           char.Level,
		Experience:      char.Experience,
		AttributePoints: char.AttributePoints,
		HP:              char.HP,
		Attack:          char.Attack,
		Defense:         char.Defense,
		CriticalChance:  char.CriticalChance,
		RankingPoints:   char.RankingPoints,
		CreatedAt:       char.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:       char.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (s *characterService) Create(userID string, req *dto.CharacterCreateRequest) (*dto.CharacterResponse, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("ID do Usuario Invalido: %w", err)
	}
	existingChar, _ := s.characterRepo.GetByName(req.Name)
	if existingChar != nil {
		return nil, fmt.Errorf("Nome do Personagem Ja Existe")
	}

	// level := req.Level
	// if level == 0 {
	// 	level = 1
	// }
	// hp := req.HP
	// if hp == 0 {
	// 	hp = 100
	// }
	// attack := req.Attack
	// if attack == 0 {
	// 	attack = 10
	// }
	// defense := req.Defense
	// if defense == 0 {
	// 	defense = 5
	// }
	// criticalChance := req.CriticalChance
	// if criticalChance == 0 {
	// 	criticalChance = 10 // Valor padrão
	//}
	
	stats := game.InitialStats[req.Class]

	character := &model.CharacterModel{
		UserID:          userUUID,
		Name:            req.Name,
		Class:           req.Class,
		Level:           1,
		Experience:      0,
		AttributePoints: 0,
		HP:              stats.HP,
		Attack:          stats.Attack,
		Defense:         stats.Defense,
		CriticalChance:  stats.CriticalChance,
	}

	createdChar, err := s.characterRepo.Create(character)
	if err != nil {
		return nil, err
	}

	return modelToDTO(createdChar), nil
}

func (s *characterService) GetByID(id string, userID string) (*dto.CharacterResponse, error) {
	character, err := s.characterRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}
	return modelToDTO(character), nil
}

func (s *characterService) GetAll(userID string) ([]*dto.CharacterResponse, error) {
	characters, err := s.characterRepo.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.CharacterResponse
	for _, char := range characters {
		responses = append(responses, modelToDTO(char))
	}
	return responses, nil
}

func (s *characterService) Update(id string, userID string, req *dto.CharacterUpdateRequest) (*dto.CharacterResponse, error) {
	character := &model.CharacterModel{
		Name: req.Name,
		// Class:          req.Class,
		// Level:          req.Level,
		// HP:             req.HP,
		// Attack:         req.Attack,
		// Defense:        req.Defense,
		// CriticalChance: req.CriticalChance,
	}

	updatedChar, err := s.characterRepo.Update(id, userID, character)
	if err != nil {
		return nil, err
	}
	return modelToDTO(updatedChar), nil
}

func (s *characterService) Delete(id string, userID string) error {
	return s.characterRepo.Delete(id, userID)
}

func (s *characterService) AddExperience(characterID string, experience int) error {
	character, err := s.characterRepo.GetByIDNoUserFilter(characterID)
	if err != nil {
		return err
	}

	totalExperience := character.Experience + experience
	levelUps := 0

	for totalExperience >= ExperienceToLevelUp {
		totalExperience -= ExperienceToLevelUp
		levelUps++
	}

	return s.characterRepo.AddExperience(characterID, totalExperience, levelUps, levelUps)
}

func (s *characterService) SpendAttributePoints(characterID string, req *dto.SpendAttributePointsRequest) error {
	character, err := s.characterRepo.GetByIDNoUserFilter(characterID)
	if err != nil {
		return err
	}

	total := req.HP + req.Attack + req.Defense + req.CriticalChance

	if total == 0 {
		return fmt.Errorf("Nenhum ponto de atributo foi gasto")
	}

	if total > character.AttributePoints {
		return fmt.Errorf("Pontos de atributo insuficientes")
	}

	return s.characterRepo.AttributePoints(characterID, req.HP, req.Attack, req.Defense, req.CriticalChance)
}
