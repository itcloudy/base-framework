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

func (m *MockUserRepository) FindUserByUserName(DB *gorm.DB, username string) (user models.UserDetail, err error) {
	ret := m.Called(username)
	if rf, ok := ret.Get(0).(func(string) models.UserDetail); ok {
		user = rf(username)
	} else {
		user = ret.Get(0).(models.UserDetail)
	}
	return
}
func (m *MockUserRepository) FindUserByID(DB *gorm.DB, id int) (user models.UserDetail, err error) {
	ret := m.Called(id)
	if rf, ok := ret.Get(0).(func(string) models.UserDetail); ok {
		user = rf(id)
	} else {
		user = ret.Get(0).(models.UserDetail)
	}
	return
}

func (m *MockUserRepository) InsertUser(DB *gorm.DB, model models.UserCreate) (user models.UserDetail, err error) {
	return
}

// 删除用户
func (m *MockUserRepository) DeleteUser(DB *gorm.DB, ids []int) (err error) {
	return
}

func (m *MockUserRepository) UpdateUserAdmin(DB *gorm.DB, id int, isAdmin bool) (err error) {
	return
}
func (m *MockUserRepository) UpdateUserActive(DB *gorm.DB, id int, isActive bool) (err error) {
	return
}

func (m *MockUserRepository) FindUserByUserNameAndPwd(DB *gorm.DB, username, pwd string) (user models.UserDetail, err error) {
	return
}

// 查询系统接口
func (m *MockUserRepository) FindAllUser(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (users []*models.UserList, total int, err error) {
	return
}
