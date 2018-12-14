// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (mo *MockUserService) GetSelf(id int) (user models.UserDetail, err error) {
	return
}

//根据用户ID查找用户
func (mo *MockUserService) GetUserByID(id int) (user models.UserDetail, err error) {
	ret := mo.Called(id)
	if rf, ok := ret.Get(0).(func(string) models.UserDetail); ok {
		user = rf(id)
	} else {
		user = ret.Get(0).(models.UserDetail)
	}
	return user, err
}

//根据用户名查找用户
func (mo *MockUserService) FindUserByUserName(DB *gorm.DB, username string) (result models.UserDetail, err error) {
	ret := mo.Called(username)
	if rf, ok := ret.Get(0).(func(string) models.UserDetail); ok {
		result = rf(username)
	} else {
		result = ret.Get(0).(models.UserDetail)
	}
	return result, err
}

// 创建用户
func (mo *MockUserService) InsertUser(DB *gorm.DB, model models.UserCreate) (result models.UserDetail, err error) {
	return
}

// 删除用户
func (mo *MockUserService) DeleteUser(DB *gorm.DB, ids []int) (err error) {
	return
}

//管理员更改
func (mo *MockUserService) UpdateUserAdmin(DB *gorm.DB, id int, isAdmin bool) (err error) {
	return
}

// 用户有效更改
func (mo *MockUserService) UpdateUserActive(DB *gorm.DB, id int, isActive bool) (err error) {
	return
}

//根据密码和用户名查询用户
func (mo *MockUserService) FindUserByUserNameAndPwd(DB *gorm.DB, username, pwd string) (model models.UserDetail, err error) {
	return
}

// 查询系统接口
func (mo *MockUserService) FindAllUser(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (users []*models.UserList, pagination conf.Pagination, err error) {
	return
}
