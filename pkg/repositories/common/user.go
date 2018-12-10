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
}

func (repo *UserRepository) FindUserByUserName(DB *gorm.DB, username string) (user models.UserDetail, err error) {
	err = DB.Where("username = ?", username).First(&user).Error
	return
}
func (repo *UserRepository) FindUserByID(DB *gorm.DB, id string) (user models.UserDetail, err error) {
	err = DB.Where("id = ?", id).First(&user).Error
	return
}
func (repo *UserRepository) InsertUser(DB *gorm.DB, create models.UserCreate) (user models.UserDetail, err error) {
	err = DB.Create(&create).Error
	if err == nil {
		return repo.FindUserByID(DB, strconv.Itoa(create.ID))
	}
	return
}
func (repo *UserRepository) UpdateUserAdmin(DB *gorm.DB, id string, isAdmin bool) (err error) {
	err = DB.Model(models.UserDetail{}).Updates(map[string]interface{}{"IsAdmin": isAdmin}).Error
	return
}
func (repo *UserRepository) UpdateUserActive(DB *gorm.DB, id string, isActive bool) (err error) {
	err = DB.Model(models.UserDetail{}).Updates(map[string]interface{}{"IsActive": isActive}).Error
	return
}

func (repo *UserRepository) FindUserByUserNameAndPwd(DB *gorm.DB, username, pwd string) (user models.UserDetail, err error) {
	err = DB.Model(models.UserDetail{}).Where("username = ? and pwd = ? and is_active = ?", username, pwd, true).First(&user).Error
	return
}
