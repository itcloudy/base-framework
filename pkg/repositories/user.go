// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
)

type IUserRepository interface {
	//根据用户名查找用户
	FindUserByUserName(username string) (models.UserDetail, error)
	//根据用户ID查找用户
	FindUserByID(id string) (models.UserDetail, error)
	// 创建用户
	InsertUser(create models.UserCreate) (models.UserDetail, error)
	//管理员更改
	UpdateUserAdmin(id string, isAdmin bool) error
	// 用户有效更改
	UpdateUserActive(id string, isActive bool) error
	//根据密码和用户名查询用户
	FindUserByUserNameAndPwd(username,pwd string) (models.UserDetail, error)
}
