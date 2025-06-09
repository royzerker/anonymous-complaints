package user

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"anonymous-complaints/internal/shared"

	"github.com/google/uuid"
)

var validRoles = []shared.RoleUser{
	shared.RoleUserAdmin,
	shared.RoleUserUser,
	shared.RoleUserGuest,
}

type User struct {
	ID        string          `json:"id"`
	Email     string          `json:"email"`
	Password  string          `json:"_"` //hash password
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Role      shared.RoleUser `json:"role"`
}

var (
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrPasswordTooShort  = errors.New("password must be at least 6 characters long")
	ErrInvalidRole       = errors.New("invalid user role")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrIsRequired        = errors.New("this field is required")
)

func NewUser(email, password, firstName, lastName string, role shared.RoleUser) (*User, error) {
	if email == "" || password == "" || firstName == "" || lastName == "" {
		return nil, ErrIsRequired
	}

	if !IsValidEmail(email) {
		return nil, ErrInvalidEmail
	}
	if len(password) < 6 {
		return nil, ErrPasswordTooShort
	}

	if !IsValidRole(role) {
		return nil, ErrInvalidRole
	}

	return &User{
		ID:        generateID(),
		Email:     strings.ToLower(email),
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      role,
	}, nil
}

func generateID() string {
	id := uuid.New().String()
	return strings.ReplaceAll(id, "-", "")
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(strings.ToLower(email))
}

func IsValidRole(role shared.RoleUser) bool {
	for _, r := range validRoles {
		if r == role {
			return true
		}
	}
	return false
}
