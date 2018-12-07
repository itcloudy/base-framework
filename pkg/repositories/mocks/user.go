// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mocks

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/mock"
)

type IUserRepository struct {
	mock.Mock
}

func (m *IUserRepository) GetUserByUserName(username string) (models.User, error) {
	ret := m.Called(username)
	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(models.User)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1

}
func (m *IUserRepository) GetUserByID(id int) (models.User, error) {
	ret := m.Called(id)
	var r0 models.User
	if rf, ok := ret.Get(0).(func(int) models.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.User)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}
