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

func (u UserRepoImpl) ViewByEmail(email string) (*model.User, error) {
	var user = model.User{}
	err := u.DB.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepo.ViewByEmail] Error execute query %v \n", err)
		return nil, fmt.Errorf("email is not exist")
	}
	return &user, nil
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
