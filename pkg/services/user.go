// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/repositories"
	"github.com/itcloudy/base-framework/tools"
)

const salt = "cloudy"

type UserService struct {
	repositories.IUserRepository
}

func (service *UserService) GetSelf(id string) (user models.UserDetail, err error) {
	return service.FindUserByID(id)
}
func (service *UserService) GetUserByID(id string) (user models.UserDetail, err error) {
	return service.FindUserByID(id)
}
func (service *UserService) GetUserByUserName(username string) (user models.UserDetail, err error) {
	return service.FindUserByUserName(username)
}
func (service *UserService) UserCreate(userCreate models.UserCreate) (user models.UserDetail, err error) {
	userCreate.ID = 0
	userCreate.Pwd = tools.SHA256(tools.StringsJoin(userCreate.Password, salt))
	return service.InsertUser(userCreate)

}
func (service *UserService) CheckUser(username, pwd string) (user models.UserDetail, err error) {
	loginPwd := tools.SHA256(tools.StringsJoin(pwd, salt))
	user, err = service.FindUserByUserNameAndPwd(username,loginPwd)
	return

}
