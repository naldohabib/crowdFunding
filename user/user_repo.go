package user

import "crowFunding/model"

type UserRepo interface {
	ViewAll() (*[]model.User, error)
	Insert(user *model.User) (*model.User, error)
	CheckMail(user *model.User) bool
	Login(user *model.User) (*model.User, error)
	FindById(ID int) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}
