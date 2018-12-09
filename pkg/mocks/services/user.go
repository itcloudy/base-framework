// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (mo *MockUserService) GetSelf(id string) (user models.UserDetail, err error) {
	return
}

func (mo *MockUserService) GetUserByID(id string) (user models.UserDetail,err  error) {
	ret := mo.Called(id)
	if rf, ok := ret.Get(0).(func(string) models.UserDetail); ok {
		user = rf(id)
	} else {
		user = ret.Get(0).(models.UserDetail)
		}

	return user, err
}
func (mo *MockUserService) GetUserByUserName(username string) (user models.UserDetail, err error) {
	return
}
func (mo *MockUserService) UserCreate(userCreate models.UserCreate) (user models.UserDetail, err error) {
	return

}
func (mo *MockUserService) CheckUser(username, pwd string) (user models.UserDetail, err error) {
	return

}
