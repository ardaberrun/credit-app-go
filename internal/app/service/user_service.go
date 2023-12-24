package service

import (
	"github/ardaberrun/credit-app-go/internal/app/model"
	"github/ardaberrun/credit-app-go/internal/app/repository"
)


type UserService struct {
	userRepository repository.IUserRepository
}

type IUserService interface {
	CreateUser(user *model.User) error
	GetUsers() ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

func InitializeUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository};
}

func (us *UserService) CreateUser(user *model.User) error {
	return us.userRepository.CreateUser(user);
}

func (us *UserService) GetUsers() ([]*model.User, error) {
	return us.userRepository.GetUsers();
}

func (us *UserService) GetUserById(id int) (*model.User, error) {
	return us.userRepository.GetUserById(id);
}

func (us *UserService) GetUserByEmail(email string) (*model.User, error) {
	return us.userRepository.GetUserByEmail(email);
}