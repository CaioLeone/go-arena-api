package repository

import (
	"database/sql"
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/model"
	"github.com/google/uuid"
)

// Interface para operacoes de Usuario
type UserRepository interface {
	CreateUser(email, name, hashedPassword string) (*model.UserModel, error)
	GetUserByEmail(email string) (*model.UserModel, error)
	GetUserByID(id string) (*model.UserModel, error)
}

// userRepository implementa UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository cria nova instancia do repositorio
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser cria novo usuario no banco
func (r *userRepository) CreateUser(email, name, hashedPassword string) (*model.UserModel, error) {
	id := uuid.New().String()

	query := `
		INSERT INTO users (id, email, name, password, created_at)
		VALUE ($1, $2, $3, $4, CURRENT_TIMESTAMP)
		RETURNING id, email, name, password, created_at
	`

	var user model.UserModel
	err := r.db.QueryRow(query, id, email, name, hashedPassword).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if err.Error() == "pq : duplicate key value violate unique constraint \"user_email_key\"" {
			return nil, fmt.Errorf("Email ja registrado")
		}
		return nil, fmt.Errorf("Erro ao criar usuario: %w", err)
	}
	return &user, nil
}

// GetUserByEmail busca usuario por email
func (r *userRepository) GetUserByEmail(email string) (*model.UserModel, error) {
	query := `
		SELECT id, email, name, password, created_at 
		FROM users 
		WHERE email = $1 
		LIMIT 1
	`
	var user model.UserModel
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Usuario nao encontrado")
	}

	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar usuario: %w", err)
	}
	return &user, nil
}

// GetUserById busca usuario por ID
func (r *userRepository) GetUserByID(id string) (*model.UserModel, error) {
	query := `
		SELECT id, email, name, password, created_at 
		FROM users 
		WHERE id = $1 
		LIMIT 1
	`

	var user model.UserModel
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return &user, nil
}
