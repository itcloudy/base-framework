// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/mocks/services"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUserController_GetUserByID(t *testing.T) {
	// 创建一个测试实例对象
	userService := new(services.MockUserService)
	id := "1"
	var (
		user   models.UserDetail
		result models.UserDetail
	)
	user.ID = 1
	user.Username = "admin"
	//设置期望结果
	userService.On("GetUserByID", id).Return(user,nil)
	userController := UserController{userService}
	//调用测试代码

	req := httptest.NewRequest("GET", "http://localhost:8080/user/1", nil)
	w := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/user/:id", userController.CtlGetUserByID)
	router.ServeHTTP(w, req)
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, user.ID, result.ID)

}
