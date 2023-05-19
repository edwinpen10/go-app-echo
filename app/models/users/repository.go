package models

import (
	"myBE/config"
)

func Save(user User) (User, error) {
	err := config.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func FindByID(id string) (User, error) {
	var user User

	err := config.DB.Where("user_uuid = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func FindByEmail(id string) (User, error) {
	var user User

	err := config.DB.Where("email = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func Update(user User) (User, error) {
	err := config.DB.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
