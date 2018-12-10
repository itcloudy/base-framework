// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mocks

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUserByID(DB *gorm.DB,id string) (user models.UserDetail, err error) {
	ret := m.Called(id)
	if rf, ok := ret.Get(0).(func(string) models.UserDetail); ok {
		user = rf(id)
	} else {
		user = ret.Get(0).(models.UserDetail)
	}
	return
}

func (m *MockUserRepository) FindUserByUserName(DB *gorm.DB,username string) (user models.UserDetail, err error) {
	return
}

func (m *MockUserRepository) InsertUser(DB *gorm.DB,create models.UserCreate) (user models.UserDetail, err error) {
	return
}
func (m *MockUserRepository) UpdateUserAdmin(DB *gorm.DB,id string, isAdmin bool) (err error) {
	return
}
func (m *MockUserRepository) UpdateUserActive(DB *gorm.DB,id string, isActive bool) (err error) {
	return
}

func (m *MockUserRepository) FindUserByUserNameAndPwd(DB *gorm.DB,username, pwd string) (user models.UserDetail, err error) {
	return
}
