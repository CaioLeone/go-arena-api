package repository

import (
	"database/sql"
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/model"
)

type CharacterRepository interface {
	Create(character *model.CharacterModel) (*model.CharacterModel, error)
	GetByID(id string, userID string) (*model.CharacterModel, error)
	GetByIDWithoutUser(id string) (*model.CharacterModel, error)
	GetAllByUserID(userID string) ([]*model.CharacterModel, error)
	GetByIDNoUserFilter(id string) (*model.CharacterModel, error)
	Update(id string, userID string, character *model.CharacterModel) (*model.CharacterModel, error)
	Delete(id string, userID string) error
	GetByName(name string) (*model.CharacterModel, error)
	AddExperience(characterID string, experience int) error
	AttributePoints(characterID string, hp int, attack int, defense int, criticalChance int) error
}

type characterRepository struct {
	db *sql.DB
}

func NewCharacterRepository(db *sql.DB) CharacterRepository {
	return &characterRepository{db: db}
}

func (r *characterRepository) Create(character *model.CharacterModel) (*model.CharacterModel, error) {
	query := `
        INSERT INTO characters (user_id, name, class, level, experience, hp, attack, defense, attribute_points, critical_chance)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING id, 
			user_id, 
			name, 
			class, 
			level, 
			experience, 
			hp, 
			attack, 
			defense, 
			attribute_points, 
			critical_chance, 
			ranking_points, 
			created_at, 
			updated_at
    `

	row := r.db.QueryRow(
		query,
		character.UserID,
		character.Name,
		character.Class,
		character.Level,
		character.Experience,
		character.HP,
		character.Attack,
		character.Defense,
		character.AttributePoints,
		character.CriticalChance,
	)

	var char model.CharacterModel
	err := row.Scan(
		&char.ID,
		&char.UserID,
		&char.Name,
		&char.Class,
		&char.Level,
		&char.Experience,
		&char.HP,
		&char.Attack,
		&char.Defense,
		&char.AttributePoints,
		&char.CriticalChance,
		&char.RankingPoints,
		&char.CreatedAt,
		&char.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Erro ao criar personagem")
		}
		return nil, fmt.Errorf("Erro ao criar personagem: %v", err)
	}
	return &char, nil
}

// func (r *characterRepository) GetByID(id string, userID string) (*model.CharacterModel, error)
func (r *characterRepository) GetByID(id string, userID string) (*model.CharacterModel, error) {
	query := `
        SELECT id, 
			user_id, 
			name, 
			class, 
			level, 
			hp, 
			experience,
			attack, 
			defense, 
			attribute_points,
			critical_chance, 
			ranking_points, 
			created_at, 
			updated_at
        FROM characters
        WHERE id = $1 AND user_id = $2
    `

	row := r.db.QueryRow(query, id, userID)
	var char model.CharacterModel
	err := row.Scan(
		&char.ID,
		&char.UserID,
		&char.Name,
		&char.Class,
		&char.Level,
		&char.HP,
		&char.Experience,
		&char.Attack,
		&char.Defense,
		&char.AttributePoints,
		&char.CriticalChance,
		&char.RankingPoints,
		&char.CreatedAt,
		&char.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Personagem não encontrado")
		}
		return nil, fmt.Errorf("Erro ao buscar personagem: %v", err)
	}
	return &char, nil

}

func (r *characterRepository) GetByIDWithoutUser(id string) (*model.CharacterModel, error) {
	query := `
        SELECT id, 
			user_id, 
			name, 
			class, 
			level, 
			experience,
			hp, 
			attack, 
			defense, 
			attribute_points,
			critical_chance, 
			ranking_points, 
			created_at, 
			updated_at
        FROM characters
        WHERE id = $1 
    `

	row := r.db.QueryRow(query, id)
	var char model.CharacterModel
	err := row.Scan(
		&char.ID,
		&char.UserID,
		&char.Name,
		&char.Class,
		&char.Level,
		&char.Experience,
		&char.HP,
		&char.Attack,
		&char.Defense,
		&char.AttributePoints,
		&char.CriticalChance,
		&char.RankingPoints,
		&char.CreatedAt,
		&char.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Personagem não encontrado")
		}
		return nil, fmt.Errorf("Erro ao buscar personagem: %v", err)
	}
	return &char, nil

}

func (r *characterRepository) GetAllByUserID(userID string) ([]*model.CharacterModel, error) {
	query := `
        SELECT id, 
			user_id, 
			name, 
			class, 
			level, 
			experience,
			hp, 
			attack, 
			defense, 
			attribute_points,
			critical_chance, 
			ranking_points, 
			created_at, 
			updated_at
        FROM characters
        WHERE user_id = $1
        ORDER BY created_at DESC
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar personagens: %v", err)
	}
	defer rows.Close()

	var characters []*model.CharacterModel
	for rows.Next() {
		var char model.CharacterModel
		err := rows.Scan(
			&char.ID,
			&char.UserID,
			&char.Name,
			&char.Class,
			&char.Level,
			&char.Experience,
			&char.HP,
			&char.Attack,
			&char.Defense,
			&char.AttributePoints,
			&char.CriticalChance,
			&char.RankingPoints,
			&char.CreatedAt,
			&char.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("Erro ao ler personagem: %v", err)
		}
		characters = append(characters, &char)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Erro ao iterar personagens: %w", err)
	}
	return characters, nil
}

func (r *characterRepository) GetByIDNoUserFilter(id string) (*model.CharacterModel, error) {
	query := `
        SELECT id, 
			user_id, 
			name, 
			class, 
			level, 
			experience,
			hp, 
			attack, 
			defense, 
			attribute_points,
			critical_chance, 
			ranking_points, 
			created_at, 
			updated_at
        FROM characters
        WHERE id = $1
    `

	row := r.db.QueryRow(query, id)

	var char model.CharacterModel
	err := row.Scan(
		&char.ID,
		&char.UserID,
		&char.Name,
		&char.Class,
		&char.Level,
		&char.Experience,
		&char.HP,
		&char.Attack,
		&char.Defense,
		&char.AttributePoints,
		&char.CriticalChance,
		&char.RankingPoints,
		&char.CreatedAt,
		&char.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Personagem não encontrado")
		}
		return nil, fmt.Errorf("Erro ao buscar personagem: %w", err)
	}

	return &char, nil
}

func (r *characterRepository) Update(id string, userID string, character *model.CharacterModel) (*model.CharacterModel, error) {
	query := `
    UPDATE characters
    SET name = COALESCE(NULLIF($1, ''), name),
        class = COALESCE(NULLIF($2, ''), class),
        level = CASE WHEN $3 > 0 THEN $3 ELSE level END,
        hp = CASE WHEN $4 > 0 THEN $4 ELSE hp END,
        attack = CASE WHEN $5 > 0 THEN $5 ELSE attack END,
        defense = CASE WHEN $6 > 0 THEN $6 ELSE defense END,
		experience = CASE WHEN $7 > 0 THEN $7 ELSE experience END,
        attribute_points = CASE WHEN $8 > 0 THEN $8 ELSE attribute_points END,
        critical_chance = CASE WHEN $9 > 0 THEN $9 ELSE critical_chance END,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = $10 AND user_id = $11
    RETURNING id, 
			user_id, 
			name, 
			class, 
			level, 
			experience, 
			hp, 
			attack, 
			defense, 
			attribute_points, 
			critical_chance, 
			ranking_points, 
			created_at, 
			updated_at
`

	row := r.db.QueryRow(
		query,
		character.Name,
		character.Class,
		character.Level,
		character.Experience,
		character.HP,
		character.Attack,
		character.Defense,
		character.AttributePoints,
		character.CriticalChance,
		id,
		userID,
	)

	var char model.CharacterModel
	err := row.Scan(
		&char.ID,
		&char.UserID,
		&char.Name,
		&char.Class,
		&char.Level,
		&char.Experience,
		&char.HP,
		&char.Attack,
		&char.Defense,
		&char.AttributePoints,
		&char.CriticalChance,
		&char.RankingPoints,
		&char.CreatedAt,
		&char.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Personagem Nao Encontrado")
		}
		return nil, fmt.Errorf("Error Ao Atualizar Personagem: %w", err)
	}
	return &char, nil
}

func (r *characterRepository) Delete(id string, userID string) error {
	query := `DELETE FROM characters WHERE id = $1 AND user_id = $2`

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("Erro ao deletar personagem: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao verificar deleção: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Personagem não encontrado")
	}
	return nil
}

func (r *characterRepository) GetByName(name string) (*model.CharacterModel, error) {
	query := `
		SELECT id, 
			user_id, 
			name, 
			class, 
			level, 
			hp, 
			attack, 
			defense, 
			experience, 
			attribute_points, 
			critical_chance, 
			ranking_points, 
			created_at, 
			updated_at
		FROM characters
		WHERE name = $1
	`

	row := r.db.QueryRow(query, name)

	var char model.CharacterModel
	err := row.Scan(
		&char.ID,
		&char.UserID,
		&char.Name,
		&char.Class,
		&char.Level,
		&char.HP,
		&char.Attack,
		&char.Defense,
		&char.Experience,
		&char.AttributePoints,
		&char.CriticalChance,
		&char.RankingPoints,
		&char.CreatedAt,
		&char.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Personagem não encontrado")
		}
		return nil, fmt.Errorf("Erro ao buscar personagem: %w", err)
	}

	return &char, nil
}

func (r *characterRepository) AddExperience(characterID string, experience int) error {
	query := `
		UPDATE characters
		SET experience = experience + $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	result, err := r.db.Exec(query, experience, characterID)
	if err != nil {
		return fmt.Errorf("Erro ao adicionar experiência: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao verificar atualização de experiência: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("Personagem não encontrado")
	}
	return nil
}

func (r *characterRepository) AttributePoints(characterID string, hp int, attack int, defense int, criticalChance int) error {
	totalSpent := hp + attack + defense + criticalChance

	query := `
		UPDATE characters 
		SET 
			hp = hp + $1,
			attack = attack + $2,
			defense = defense + $3,
			critical_chance = critical_chance + $4,
			attribute_points = attribute_points - $5,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $6
	`

	result, err := r.db.Exec(query, hp, attack, defense, criticalChance, totalSpent, characterID)

	if err != nil {
		return fmt.Errorf("Erro ao gastar pontos de atributo: %w", err)

	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao verificar atualização de pontos de atributo: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("Personagem não encontrado")
	}
	return nil
}
