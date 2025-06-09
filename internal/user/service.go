package user

import (
	"anonymous-complaints/internal/shared"
	"anonymous-complaints/pkg/hash"
	"anonymous-complaints/pkg/jwt"
	"errors"
	"time"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (a *UserService) Register(email, password string, role shared.RoleUser) error {
	existingUser, _ := a.userRepo.FindByEmail(email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	if !IsValidRole(role) {
		return errors.New("invalid role")
	}

	hashed, err := hash.HashPassword(password)
	if err != nil {
		return err
	}

	user := &User{
		Email:    email,
		Password: hashed,
		Role:     role,
	}

	return a.userRepo.Create(user)
}

func (a *UserService) Login(email, password string) (string, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generar JWT con 24h expiración
	token, err := jwt.GenerateToken(user.ID, user.Role, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}
