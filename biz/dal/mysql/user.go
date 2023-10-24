package mysql

import (
	"github.com/pkg/errors"
	"hertz/demo/biz/model/mysql"
	"hertz/demo/pkg/utils"
)

func GetUserList(cursor, num int64) ([]mysql.User, error) {
	var userList []mysql.User
	if err := DB.Where("id > ?", cursor).Limit(int(num)).Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func CreateUser(username, passwd string) error {
	pwdHash, err := utils.GetHashPassword(passwd)
	if err != nil {
		return errors.WithMessage(err, "get hash password err")
	}
	if err = DB.Create(&mysql.User{
		Username: username,
		Password: pwdHash,
	}).Error; err != nil {
		return errors.WithMessage(err, "create user err")
	}
	return nil
}

func DeleteUser(userID int64) error {
	var user mysql.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.WithMessage(err, "no such user")
	}
	if err := DB.Where("id = ?", userID).Delete(&user).Error; err != nil {
		return errors.WithMessage(err, "delete user err")
	}
	return nil
}

func DisableUser(userID int64) error {
	var user mysql.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.WithMessage(err, "no such user")
	}
	if err := DB.Where("id = ?", userID).Update("is_disabled", true).Error; err != nil {
		return errors.WithMessage(err, "delete user err")
	}
	return nil
}

func GetUserByName(username string) (mysql.User, error) {
	var user mysql.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserBySub(sub int64) (mysql.User, error) {
	var user mysql.User
	if err := DB.Where("id = ?", sub).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
