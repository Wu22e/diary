package service

import "crud_server/repository"

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepsitory *repository.UserRepository) *User {
	return &User{
		userRepository: userRepsitory,
	}
}
