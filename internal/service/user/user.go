// Package user. UserService methods
package user

import "e-commerce-system/internal/model/user"

type Repository interface {
	Add(user *user.User) error
	Delete(id int) error
	GetById(id int) (error, *user.User)
	List() []*user.User
}

type Service struct {
	userRepository Repository
}

func CreateUserService(userRepository Repository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) CreateUser(id int, firstName, lastName string, age int8) error {
	return s.userRepository.Add(user.CreateUser(id, firstName, lastName, age))
}

func (s *Service) GetById(id int) (error, *user.User) {
	// some instructions ...
	return s.userRepository.GetById(id)
}

func (s *Service) UserList() []*user.User {
	return s.userRepository.List()
}

func (s *Service) DeleteUserById(id int) error {
	return s.userRepository.Delete(id)
}
