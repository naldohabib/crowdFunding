package service

import (
	"crowFunding/model"
	"crowFunding/user"
)

type UserServiceImpl struct {
	userRepo user.UserRepo
}

func (u UserServiceImpl) ViewByEmail(email string) (*model.User, error) {
	return u.userRepo.ViewByEmail(email)
}

func (u UserServiceImpl) Insert(user *model.User) (*model.User, error) {
	return u.userRepo.Insert(user)
}

func (u UserServiceImpl) ViewAll() (*[]model.User, error) {
	return u.userRepo.ViewAll()
}

func CreateUserService(userRepo user.UserRepo) user.UserService {
	return &UserServiceImpl{userRepo}
}
