// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/services/mocks"
	"testing"
)

func TestUserController_GetUserByID(t *testing.T) {
	userService :=new(mocks.IUserService)
	id :="1"
	var user models.UserDetail
	user.ID = 1
	userService.On("GetUserByID",id).Return(user,nil)
	userController :=UserController{UserService:userService}
}
