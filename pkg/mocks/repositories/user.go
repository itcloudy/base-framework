// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mocks

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUserByID(id string) (user models.UserDetail, err error) {
	ret := m.Called(id)
	if rf, ok := ret.Get(0).(func(string) models.UserDetail); ok {
		user = rf(id)
	} else {
		user = ret.Get(0).(models.UserDetail)
	}
	return
}

func (m *MockUserRepository) FindUserByUserName(username string) (user models.UserDetail, err error) {
	return
}

func (m *MockUserRepository) InsertUser(create models.UserCreate) (user models.UserDetail, err error) {
	return
}
func (m *MockUserRepository) UpdateUserAdmin(id string, isAdmin bool) (err error) {
	return
}
func (m *MockUserRepository) UpdateUserActive(id string, isActive bool) (err error) {
	return
}

func (m *MockUserRepository) FindUserByUserNameAndPwd(username, pwd string) (user models.UserDetail, err error) {
	return
}
