package service

import (
	"fmt"

	"github.com/caioLeone/go-arena-api/internal/auth"
	"github.com/caioLeone/go-arena-api/internal/config"
	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthService interface para autenticacao
type AuthService interface {
	Register(req *dto.UserCreateRequest) (*dto.UserResponse, *dto.LoginResponse, error)
	Login(req *dto.UserLoginRequest) (*dto.LoginResponse, error)
	ValidateToken(token string) (string, error)
}

// authService implementa AuthService
type authService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

// NewAuthService criar nova instancia do servico
func NewAuthService(userRepo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		config:   cfg,
	}
}

// Register registra novo usuario
func (s *authService) Register(req *dto.UserCreateRequest) (*dto.UserResponse, *dto.LoginResponse, error) {
	//validar se email ja existe
	_, err := s.userRepo.GetUserByEmail(req.Email)
	if err == nil {
		return nil, nil, fmt.Errorf("Email ja registrado")
	}

	//Hash da senha com bcrypt
	hasehdPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao Hashear senha: %w", err)
	}

	//Criar usuario(nome = email por enquanto, sera atualizado depois)
	user, err := s.userRepo.CreateUser(req.Email, req.Email, string(hasehdPassword))
	if err != nil {
		return nil, nil, err
	}

	accessToken, err := auth.GenerateAccessToken(user.ID.String(), user.Email, s.config)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao Gerar Access Token: %w", err)
	}

	refreshToken := auth.GenerateRefreshtToken()

	//Response
	userResp := &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	loginResp := &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *userResp,
	}
	return userResp, loginResp, nil
}

// Login Faz Login do Usuario
func (s *authService) Login(req *dto.UserLoginRequest) (*dto.LoginResponse, error) {
	//buscar usuario por email
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("Credenciais Invalidas")
	}

	//validar senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("Credenciais invalidas")
	}

	//Gerar Tokens
	accessToken, err := auth.GenerateAccessToken(user.ID.String(), user.Email, s.config)
	if err != nil {
		return nil, fmt.Errorf("Erro ao Gerar Access Token: %w", err)
	}

	refreshToken := auth.GenerateRefreshtToken()

	// Response
	userResp := &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	loginResp := &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *userResp,
	}

	return loginResp, nil
}

func (s *authService) ValidateToken(token string) (string, error) {
	claims, err := auth.ValidateToken(token, s.config)
	if err != nil {
		return "", fmt.Errorf("Token Invalido: %w", err)
	}

	return claims.UserID, nil
}
