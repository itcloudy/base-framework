// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	*gorm.DB
}

func (repo *UserRepository) FindUserByUserName(username string) (user models.User, err error) {
	err = repo.Where("username = ?", username).First(&user).Error
	return
}
func (repo *UserRepository) FindUserByID(id string) (user models.User, err error) {
	err = repo.Where("id = ?", id).First(&user).Error
	return
}
