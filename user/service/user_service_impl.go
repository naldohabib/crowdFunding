package service

import (
	"crowFunding/model"
	"crowFunding/user"
	"crowFunding/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepo user.UserRepo
}

func (u UserServiceImpl) SaveAvatar(ID int, fileLocation string) (*model.User, error) {
	user, err := u.userRepo.FindById(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updateUser, err := u.userRepo.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (u UserServiceImpl) Login(user *model.User) (*model.User, error) {
	model, err := u.userRepo.Login(user)
	if err != nil {
		return nil, errors.New("Email not registered")
	}

	err = utils.VerifyPassword(model.Password, user.Password)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New("Email invalid password")
	}

	return model, nil
}

func (u UserServiceImpl) CheckMail(user *model.User) bool {
	return u.userRepo.CheckMail(user)
}

func (u UserServiceImpl) Insert(user *model.User) (*model.User, error) {
	chechMail := u.userRepo.CheckMail(user)

	if !chechMail {
		fmt.Println("[UserService.Insert] Email already used!")
		return nil, errors.New("Opps.. sorry email already used")
	}
	return u.userRepo.Insert(user)
}

func (u UserServiceImpl) ViewAll() (*[]model.User, error) {
	return u.userRepo.ViewAll()
}

func CreateUserService(userRepo user.UserRepo) user.UserService {
	return &UserServiceImpl{userRepo}
}
