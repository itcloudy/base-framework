// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/mocks/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestUserService_GetUserByID(t *testing.T) {
	userRepository := new(mocks.MockUserRepository)
	idStr := "101"
	idInt, _ := strconv.Atoi(idStr)
	var user models.UserDetail
	user.ID = idInt
	userRepository.On("FindUserByID", idStr).Return(user, nil)
	userService := UserService{IUserRepository: userRepository}
	result, _ := userService.GetUserByID(idStr)
	assert.Equal(t, user.ID, result.ID)
}
func TestUserService_GetUserByUserName(t *testing.T) {
	userRepository := new(mocks.MockUserRepository)
	name := "cloudy"
	var user models.UserDetail
	user.Username = name
	userRepository.On("FindUserByUserName", name).Return(user, nil)
	// 判断测试结果
	assert.Equal(t, name, user.Username)
}
