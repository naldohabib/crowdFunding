package user

import "crowFunding/model"

type UserService interface {
	ViewAll()(*[]model.User, error)
	Insert(user *model.User)(*model.User, error)
	ViewByEmail(email string)(*model.User, error)
}
