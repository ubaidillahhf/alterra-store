package database

import (
	"alterra_store/configs"
	"alterra_store/models/users"
)

func GetDataUserAll() (dataResult []users.Users, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}
