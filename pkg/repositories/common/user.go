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
func (repo *UserRepository) InsertUser(DB *gorm.DB, model models.UserCreate) (user models.UserDetail, err error) {
	model.ID = 0
	err = DB.Create(&model).Error
	if err == nil {
		return repo.FindUserByID(DB, strconv.Itoa(model.ID))
	}
	return
}
func (repo *UserRepository) DeleteUser(DB *gorm.DB, ids []string) (err error) {
	return DB.Where("id IN (?)", ids).Delete(models.UserDetail{}).Error
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
	err = DB.Where("username = ? and pwd = ? and is_active = ?", username, pwd, true).First(&user).Error
	return
}

// 查询用户
func (repo *UserRepository) FindAllUser(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (results []*models.UserList, total int, err error) {
	if len(order) == 0 {
		order = "id desc"
	}
	db := DB.Model(&models.UserList{}).Order(order)
	if len(query) > 0 {
		db = db.Where(query, queryArgs[:])
	}
	db.Count(&total)
	err = db.Offset((page - 1) * size).Limit(size).Find(&results).Error
	return
}
