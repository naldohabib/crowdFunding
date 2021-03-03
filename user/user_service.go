package user

import "crowFunding/model"

type UserService interface {
	ViewAll() (*[]model.User, error)
	Insert(user *model.User) (*model.User, error)
	CheckMail(user *model.User) bool
	Login(user *model.User) (*model.User, error)
	SaveAvatar(ID int, fileLocation string) (*model.User, error)
}
