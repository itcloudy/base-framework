// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

type UserRepository struct {
	*gorm.DB
}

func (repo *UserRepository) FindUserByUserName(username string) (user models.UserDetail, err error) {
	err = repo.Where("username = ?", username).First(&user).Error
	return
}
func (repo *UserRepository) FindUserByID(id string) (user models.UserDetail, err error) {
	err = repo.Where("id = ?", id).First(&user).Error
	return
}
func (repo *UserRepository) InsertUser(create models.UserCreate) (user models.UserDetail, err error) {
	err = repo.Create(&create).Error
	if err == nil {
		return repo.FindUserByID(strconv.Itoa(create.ID))
	}
	return
}
func (repo *UserRepository) UpdateUserAdmin(id string, isAdmin bool) (err error) {
	err = repo.Model(models.User{}).Updates(map[string]interface{}{"IsAdmin": isAdmin}).Error
	return
}
func (repo *UserRepository) UpdateUserActive(id string, isActive bool) (err error) {
	err = repo.Model(models.User{}).Updates(map[string]interface{}{"IsActive": isActive}).Error
	return
}

func (repo *UserRepository) FindUserByUserNameAndPwd(username, pwd string) (user models.UserDetail, err error) {
	err = repo.Model(models.User{}).Where("username = ? and pwd = ? and is_active = ?", username, pwd, true).First(&user).Error

	return
}
