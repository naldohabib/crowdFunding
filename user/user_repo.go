package user

import "crowFunding/model"

type UserRepo interface {
	ViewAll()(*[]model.User, error)
	Insert(user *model.User) (*model.User, error)
	ViewByEmail(email string)(*model.User, error)
}
