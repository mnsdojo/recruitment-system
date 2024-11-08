package services

import (
	"errors"

	"github.com/mnsdojo/recruitment-system/models"
	"github.com/mnsdojo/recruitment-system/repository"
)

type UserService interface {
	RegisterUser(user *models.User) (*models.User, error)
	AuthenticateUser(email, password string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) RegisterUser(user *models.User) (*models.User, error) {
	// Check if the email is already registered
	existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Create the user
	return s.userRepo.CreateUser(user)
}

func (s *userService) AuthenticateUser(email, password string) (*models.User, error) {
	return s.userRepo.VerifyUserCredentials(email, password)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}
