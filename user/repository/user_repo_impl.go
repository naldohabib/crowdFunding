package repository

import (
	"crowFunding/model"
	"crowFunding/user"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func (u UserRepoImpl) Update(user *model.User) (*model.User, error) {
	err := u.DB.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) FindById(ID int) (*model.User, error) {
	user :=  new(model.User)

	err := u.DB.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (u UserRepoImpl) Login(user *model.User) (*model.User, error) {
	users := new(model.User)
	err := u.DB.Table("users").Where("email = ?", user.Email).Take(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (u UserRepoImpl) CheckMail(user *model.User) bool {
	users := new(model.UserFormatter)

	err := u.DB.Raw("SELECT * FROM \"users\" WHERE email = ? LIMIT 1", user.Email).Scan(users).Error

	if err != nil {
		return true
	}

	return false
}

func (u UserRepoImpl) Insert(user *model.User) (*model.User, error) {
	err := u.DB.Save(&user).Error

	if err != nil {
		fmt.Printf("[UserRepo.Insert] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data user")
	}

	return user, nil
}

func (u UserRepoImpl) ViewAll() (*[]model.User, error) {
	var user []model.User

	err := u.DB.Find(&user).Error

	if err != nil {
		fmt.Printf("[UserRepo.ViewAll] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed read all data user")
	}
	return &user, nil
}

func CreateUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}
