// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/repositories"
)

type UserService struct {
	repositories.IUserRepository
}

func (service *UserService) GetSelf(id string) (user models.User, err error) {
	return service.FindUserByID(id)
}
func (service *UserService) GetUserByID(id string) (user models.User, err error) {
	return service.FindUserByID(id)
}
func (service *UserService) GetUserByUserName(username string) (user models.User, err error) {
	return service.FindUserByUserName(username)
}
