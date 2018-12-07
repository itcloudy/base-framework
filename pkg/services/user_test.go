// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/repositories/mocks"
	"testing"
)

func TestUserService_GetUserByID(t *testing.T) {
	userRepository := new(mocks.IUserRepository)
	id := 101
	user1 := models.User{}
	user1.ID = id
	user1.Username = "cloudy"
	userRepository.On("GetUserByID", id).Return(user1, nil)
	if user1.ID != id {
		t.Errorf("service GetUserByID failed, excepted: %d, got: %d", id, user1.ID)
	}
	//userService := UserService{userRepository}
}
func TestUserService_GetUserByUserName(t *testing.T) {
	userRepository := new(mocks.IUserRepository)
	name := "cloudy"
	user1 := models.User{}
	user1.Username = name
	userRepository.On("GetUserByUserName", name).Return(user1, nil)
	if user1.Username != name {
		t.Errorf("service GetUserByUserName failed, excepted: %s, got: %s", name, user1.Username)
	}
}
