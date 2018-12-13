// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/tools"
	"github.com/jinzhu/gorm"
)

const salt = "cloudy"

type UserService struct {
	DB *gorm.DB
	repositories.IUserRepository
}

func (service *UserService) ServiceGetSelf(id string) (user models.UserDetail, err error) {
	return service.FindUserByID(service.DB, id)
}
func (service *UserService) ServiceGetUserByID(id string) (user models.UserDetail, err error) {
	return service.FindUserByID(service.DB, id)
}
func (service *UserService) ServiceGetUserByUserName(username string) (user models.UserDetail, err error) {
	return service.FindUserByUserName(service.DB, username)
}
func (service *UserService) ServiceUserCreate(model models.UserCreate) (user models.UserDetail, err error) {
	model.ID = 0
	model.Pwd = tools.SHA256(tools.StringsJoin(model.Password, salt))
	return service.InsertUser(service.DB, model)
}
func (service *UserService) ServiceUserDelete(ids []string) (err error) {
	return service.DeleteUser(service.DB, ids)
}
func (service *UserService) ServiceCheckUser(username, pwd string) (user models.UserDetail, err error) {
	loginPwd := tools.SHA256(tools.StringsJoin(pwd, salt))
	user, err = service.FindUserByUserNameAndPwd(service.DB, username, loginPwd)
	return
}
func (service *UserService) ServiceGetAllUser(offset, limit int, order string, query string, queryArgs ...interface{}) (users []*models.UserList, count int, err error) {
	return service.FindAllUser(service.DB, offset, limit, order, query, queryArgs)
}
