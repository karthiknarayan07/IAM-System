package service

import (
	"github.com/karthiknarayan07/IAM-System/db/models"
	"github.com/karthiknarayan07/IAM-System/domain"
	"github.com/karthiknarayan07/IAM-System/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(username, email string) (*domain.User, error) {
	user := &domain.User{
		ID:       uuid.New().String(),
		Username: username,
		Email:    email,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Convert to DB model and save
	dbUser := &models.User{
		ID:       uuid.MustParse(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}

	if err := s.repo.Create(dbUser); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	dbUser, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       dbUser.ID.String(),
		Username: dbUser.Username,
		Email:    dbUser.Email,
	}, nil
}
