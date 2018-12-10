// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"fmt"
	"github.com/itcloudy/base-framework/pkg/mocks/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestUserService_GetUserByID(t *testing.T) {
	userRepository := new(mocks.MockUserRepository)
	var err error
	idStr := "101"
	idInt, _ := strconv.Atoi(idStr)
	var user models.UserDetail
	userRepository.On("GetUserByID", idStr).Return(user, nil)
	user, err = userRepository.FindUserByID(idStr)
	fmt.Printf("========%+v\n",err)
	assert.Equal(t, user.ID, idInt)
}
func TestUserService_GetUserByUserName(t *testing.T) {
	userRepository := new(mocks.MockUserRepository)
	name := "cloudy"
	var user1 models.User
	userRepository.On("GetUserByUserName", name).Return(user1, nil)
	if user1.Username != name {
		t.Errorf("service GetUserByUserName failed, excepted: %s, got: %s", name, user1.Username)
	}
	// 判断测试结果
	assert.Equal(t, name, user1.Username)
}
