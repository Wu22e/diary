package service

import (
	"crud_server/repository"
	"crud_server/types"
)

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepsitory *repository.UserRepository) *User {
	return &User{
		userRepository: userRepsitory,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(name string, newAge int64) error {
	return u.userRepository.Update(name, newAge)
}

func (u *User) Delete(user *types.User) error {
	return u.userRepository.Delete(user.Name)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
