// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	//根据用户名查找用户
	FindUserByUserName(DB *gorm.DB, username string) (models.UserDetail, error)
	//根据用户ID查找用户
	FindUserByID(DB *gorm.DB, id int) (models.UserDetail, error)
	// 创建用户
	InsertUser(DB *gorm.DB, model models.UserCreate) (models.UserDetail, error)
	// 删除用户
	DeleteUser(DB *gorm.DB, ids []int) error
	//管理员更改
	UpdateUserAdmin(DB *gorm.DB, id int, isAdmin bool) error
	// 用户有效更改
	UpdateUserActive(DB *gorm.DB, id int, isActive bool) error
	//根据密码和用户名查询用户
	FindUserByUserNameAndPwd(DB *gorm.DB, username, pwd string) (models.UserDetail, error)
	// 查询用户
	FindAllUser(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (users []*models.UserList, total int, err error)
}
