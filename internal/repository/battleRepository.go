package repository

import (
	"database/sql"
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/model"
)

type BattleRepository interface {
	Create(battle *model.BattleModel) (*model.BattleModel, error)
	GetByID(id string) (*model.BattleModel, error)
	GetHistoryByUserID(userID string, limit int, offset int) ([]*model.BattleModel, error)
	UpdateCharacterRanking(characterID string, points int) error
	GetCharacterRanking(characterID string) (int, error)
}

type battleRepository struct {
	db *sql.DB
}

func NewBattleRepository(db *sql.DB) BattleRepository {
	return &battleRepository{db: db}
}

func (r *battleRepository) Create(battle *model.BattleModel) (*model.BattleModel, error) {
	query := `
        INSERT INTO battles (attacker_id, attacker_name, defender_id, defender_name, winner_id, winner_name, damage_dealt, attacker_hp_final, defender_hp_final, rounds_count, rounds_data)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id, attacker_id, attacker_name, defender_id, defender_name, winner_id, winner_name, damage_dealt, attacker_hp_final, defender_hp_final, rounds_count, rounds_data, created_at
    `

	row := r.db.QueryRow(
		query,
		battle.AttackerID,
		battle.AttackerName,
		battle.DefenderID,
		battle.DefenderName,
		battle.WinnerID,
		battle.WinnerName,
		battle.DamageDealt,
		battle.AttackerHPFinal,
		battle.DefenderHPFinal,
		battle.RoundsCount,
		battle.RoundsData,
	)

	var createdBattle model.BattleModel
	err := row.Scan(
		&createdBattle.ID,
		&createdBattle.AttackerID,
		&createdBattle.AttackerName,
		&createdBattle.DefenderID,
		&createdBattle.DefenderName,
		&createdBattle.WinnerID,
		&createdBattle.WinnerName,
		&createdBattle.DamageDealt,
		&createdBattle.AttackerHPFinal,
		&createdBattle.DefenderHPFinal,
		&createdBattle.RoundsCount,
		&createdBattle.RoundsData,
		&createdBattle.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("Erro ao criar batalha: %w", err)
	}

	return &createdBattle, nil
}

func (r *battleRepository) GetByID(id string) (*model.BattleModel, error) {
	query := `
        SELECT id, attacker_id, attacker_name, defender_id, defender_name, winner_id, winner_name, damage_dealt, attacker_hp_final, defender_hp_final, rounds_count, rounds_data, created_at
        FROM battles
        WHERE id = $1
    `

	row := r.db.QueryRow(query, id)

	var battle model.BattleModel
	err := row.Scan(
		&battle.ID,
		&battle.AttackerID,
		&battle.AttackerName,
		&battle.DefenderID,
		&battle.DefenderName,
		&battle.WinnerID,
		&battle.WinnerName,
		&battle.DamageDealt,
		&battle.AttackerHPFinal,
		&battle.DefenderHPFinal,
		&battle.RoundsCount,
		&battle.RoundsData,
		&battle.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Batalha não encontrada")
		}
		return nil, fmt.Errorf("Erro ao buscar batalha: %w", err)
	}

	return &battle, nil
}

func (r *battleRepository) GetHistoryByUserID(userID string, limit int, offset int) ([]*model.BattleModel, error) {
	query := `
        SELECT 
			b.id,  
			b.attacker_id,  
			b.attacker_name,  
			b.defender_id,  
			b.defender_name,  
			b.winner_id, 
			b.winner_name, 
			b.damage_dealt, 
			b.attacker_hp_final, 
			b.defender_hp_final, 
			b.rounds_count, 
			b.rounds_data, 
			b.created_at
        FROM battles b
		JOIN 
			characters attacker ON attacker.id = b.attacker_id
		JOIN 
			characters defender ON defender.id = b.defender_id
        WHERE 
			attacker.user_id = $1 
			OR defender.user_id = $1
        ORDER BY b.created_at DESC
        LIMIT $2 OFFSET $3
    `

	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar histórico de batalhas: %w", err)
	}
	defer rows.Close()

	var battles []*model.BattleModel
	for rows.Next() {
		var battle model.BattleModel
		err := rows.Scan(
			&battle.ID,
			&battle.AttackerID,
			&battle.AttackerName,
			&battle.DefenderID,
			&battle.DefenderName,
			&battle.WinnerID,
			&battle.WinnerName,
			&battle.DamageDealt,
			&battle.AttackerHPFinal,
			&battle.DefenderHPFinal,
			&battle.RoundsCount,
			&battle.RoundsData,
			&battle.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("Erro ao ler batalha: %w", err)
		}

		battles = append(battles, &battle)
	}

	return battles, nil
}

func (r *battleRepository) UpdateCharacterRanking(characterID string, points int) error {
	query := `UPDATE characters SET ranking_points = ranking_points + $1 WHERE id = $2`

	_, err := r.db.Exec(query, points, characterID)
	if err != nil {
		return fmt.Errorf("Erro ao atualizar ranking: %w", err)
	}

	return nil
}

func (r *battleRepository) GetCharacterRanking(characterID string) (int, error) {
	query := `SELECT ranking_points FROM characters WHERE id = $1`

	var ranking int
	err := r.db.QueryRow(query, characterID).Scan(&ranking)
	if err != nil {
		return 0, fmt.Errorf("Erro ao buscar ranking: %w", err)
	}

	return ranking, nil
}
